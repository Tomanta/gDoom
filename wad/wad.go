package wad

import (
	"fmt"
	"slices"
)

type Wad struct {
	Header    Header
	Directory Directory
	// Levels    []Level
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
	dir, err := NewDirectoryFromBytes(buf[dirStart:endPos], numLumps)
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
	for i, e := range dir.Entries {
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
	for _, l := range levelList {
		fmt.Printf("Level: %v\n", l)
	}

	w := Wad{
		Header:    header,
		Directory: dir,
	}

	return w, nil
}

func isLevel(name string) bool {
	if len(name) == 4 && name[0] == 'E' && name[2] == 'M' {
		return true
	}
	if len(name) == 5 && name[0:3] == "MAP" {
		return true
	}
	return false
}

func isLevelLump(name string) bool {
	var levelLumps = []string{"THINGS", "LINEDEFS", "SIDEDEFS", "VERTEXES", "SEGS", "SSECTORS", "NODES", "SECTORS", "REJECT", "BLOCKMAP"}

	return slices.Contains(levelLumps, name)
}
