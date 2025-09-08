package wad

import "fmt"

type Sector struct {
	FloorHeight    int16
	CeilingHeight  int16
	LightLevel     int16
	SpecialType    int16
	TagID          int16
	FloorTexture   string
	CeilingTexture string
}

func readSectorFromBuffer(buf []byte) (Sector, error) {
	floorHeight, _ := Int16FromBytes(buf[0:2])
	ceilingHeight, _ := Int16FromBytes(buf[2:4])
	floorTexture := StringFromBytes(buf[4:12])
	ceilingTexture := StringFromBytes(buf[12:20])
	lightLevel, _ := Int16FromBytes(buf[20:22])
	specialType, _ := Int16FromBytes(buf[22:24])
	tagID, _ := Int16FromBytes(buf[24:26])

	s := Sector{
		FloorHeight:    floorHeight,
		CeilingHeight:  ceilingHeight,
		LightLevel:     lightLevel,
		SpecialType:    specialType,
		TagID:          tagID,
		FloorTexture:   floorTexture,
		CeilingTexture: ceilingTexture,
	}
	return s, nil
}

func NewSectorsFromBytes(buf []byte, numEntries int32) ([]Sector, error) {
	totalSectorSize := numEntries * LUMP_SIZE_SECTOR
	if (int32)(len(buf)) != totalSectorSize {
		return []Sector{}, fmt.Errorf("expected buffer size %d, actual size %d", totalSectorSize, len(buf))
	}

	var sectors []Sector

	for i := range numEntries {
		start := i * LUMP_SIZE_SECTOR
		end := start + LUMP_SIZE_SECTOR
		entry, err := readSectorFromBuffer(buf[start:end])
		if err != nil {
			return []Sector{}, fmt.Errorf("error creating sectors: %v", err)
		}
		sectors = append(sectors, entry)

	}
	return sectors, nil
}
