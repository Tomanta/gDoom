package wad

import "fmt"

type Sidedef struct {
	TextureOffsetX int16
	TextureOffsetY int16
	SectorID       int16
	UpperTexture   string
	LowerTexture   string
	MiddleTexture  string
}

func readSidedefFromBuffer(buf []byte) (Sidedef, error) {
	textureOffsetX, _ := Int16FromBytes(buf[0:2])
	textureOffsetY, _ := Int16FromBytes(buf[2:4])
	upperTexture, _ := StringFromBytes(buf[4:12], 8)
	lowerTexture, _ := StringFromBytes(buf[12:20], 8)
	middleTexture, _ := StringFromBytes(buf[20:28], 8)
	sectorID, _ := Int16FromBytes(buf[28:30])

	sd := Sidedef{
		TextureOffsetX: textureOffsetX,
		TextureOffsetY: textureOffsetY,
		SectorID:       sectorID,
		UpperTexture:   upperTexture,
		LowerTexture:   lowerTexture,
		MiddleTexture:  middleTexture,
	}
	return sd, nil
}

func NewSidedefsFromBytes(buf []byte, numEntries int32) ([]Sidedef, error) {
	var entrySize int32 = 30
	if (int32)(len(buf)) != numEntries*entrySize {
		return []Sidedef{}, fmt.Errorf("expected buffer size %d (%d * 16 bytes), actual size %d", numEntries*entrySize, numEntries, len(buf))
	}

	var sidedefs []Sidedef

	for i := range numEntries {
		start := i * entrySize
		end := start + entrySize
		entry, err := readSidedefFromBuffer(buf[start:end])
		if err != nil {
			return []Sidedef{}, fmt.Errorf("error creating directory: %v", err)
		}
		sidedefs = append(sidedefs, entry)

	}
	return sidedefs, nil
}
