package wad

import (
	"encoding/binary"
	"fmt"
)

type Thing struct {
	XPos   int16
	YPos   int16
	Angle  int16
	TypeID int16
	Flags  int16
}

const (
	TH_SKILL_1_2 int16 = 1 << iota
	TH_SKILL_3
	TH_SKILL_4_5
	TH_AMBUSH
	TH_MULTI_ONLY
)

func readThingFromBuffer(buf []byte) (Thing, error) {
	var t Thing
	_, err := binary.Decode(buf, binary.LittleEndian, &t)
	if err != nil {
		return Thing{}, fmt.Errorf("error decoding thing: %v", err)
	}

	return t, nil
}

func NewThingsFromBytes(buf []byte, numEntries int32) ([]Thing, error) {
	totalThingSize := numEntries * LUMP_SIZE_THING
	if (int32)(len(buf)) != totalThingSize {
		return nil, fmt.Errorf("invalid buffer length; expected %d, got %d", totalThingSize, len(buf))
	}
	var things []Thing

	for i := range numEntries {
		start := i * LUMP_SIZE_THING
		end := start + LUMP_SIZE_THING
		thing, err := readThingFromBuffer(buf[start:end])
		if err != nil {
			return nil, fmt.Errorf("error creating things: %v", err)
		}
		things = append(things, thing)

	}
	return things, nil

}
