package wad

import (
	"fmt"
)

// Define the size of lumps with a constant number of bytes
const (
	LUMP_SIZE_DIRECTORY_ENTRY = 16
	LUMP_SIZE_HEADER          = 12
	LUMP_SIZE_LINEDEF         = 14
	LUMP_SIZE_SECTOR          = 26
	LUMP_SIZE_SIDEDEF         = 30
	LUMP_SIZE_THING           = 10
	LUMP_SIZE_VERTEX          = 4
)

type Wad struct {
	Header    Header
	Directory []DirEntry
	Levels    []Level
}

func NewWadFromBytes(buf []byte) (Wad, error) {
	// load header
	header, err := NewHeaderFromBytes(buf[0:12])
	if err != nil {
		return Wad{}, err
	}

	// header points us to the lump directory, load that
	dirStart := header.DirectoryPos
	numLumps := header.NumLumps
	endPos := dirStart + (numLumps * 16)
	directory, err := NewDirectoryFromBytes(buf[dirStart:endPos], numLumps)
	if err != nil {
		return Wad{}, err
	}

	type levelInfo struct {
		startIndex int
		endIndex   int
		name       string
	}

	levelList := []levelInfo{}
	isReadingLevel := false
	curLevelInfo := levelInfo{}

	// now that we have that, we can start loading level data
	for i, e := range directory {
		if isLevelLump(e.Name) {
			if !isReadingLevel {
				return Wad{}, fmt.Errorf("error reading wad, lump %s outside level definition", e.Name)
			}
			continue
		}

		if isLevel(e.Name) {
			if isReadingLevel {
				curLevelInfo.endIndex = i
				levelList = append(levelList, curLevelInfo)
			}

			isReadingLevel = true
			curLevelInfo = levelInfo{name: e.Name, startIndex: i}
			continue
		}

		if !isLevel(e.Name) && !isLevelLump(e.Name) && isReadingLevel {
			curLevelInfo.endIndex = i
			levelList = append(levelList, curLevelInfo)
			isReadingLevel = false
			continue
		}
	}

	// Loop through level list to load all the levels
	levels := []Level{}
	for _, l := range levelList {
		levelData, _ := NewLevelFromBuffer(buf, directory[l.startIndex:l.endIndex])
		levels = append(levels, levelData)
	}

	w := Wad{
		Header:    header,
		Directory: directory,
		Levels:    levels,
	}

	return w, nil
}
