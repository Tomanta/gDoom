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
	upperTexture := StringFromBytes(buf[4:12])
	lowerTexture := StringFromBytes(buf[12:20])
	middleTexture := StringFromBytes(buf[20:28])
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
	totalSidedefSize := numEntries * LUMP_SIZE_SIDEDEF
	if (int32)(len(buf)) != totalSidedefSize {
		return []Sidedef{}, fmt.Errorf("expected buffer size %d, actual size %d", totalSidedefSize, len(buf))
	}

	var sidedefs []Sidedef

	for i := range numEntries {
		start := i * LUMP_SIZE_SIDEDEF
		end := start + LUMP_SIZE_SIDEDEF
		entry, err := readSidedefFromBuffer(buf[start:end])
		if err != nil {
			return []Sidedef{}, fmt.Errorf("error creating sidedefs: %v", err)
		}
		sidedefs = append(sidedefs, entry)

	}
	return sidedefs, nil
}
