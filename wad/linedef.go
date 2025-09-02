package wad

import (
	"encoding/binary"
	"fmt"
)

const (
	LD_IMPASSIBLE     int16 = 0
	LD_BLOCK_MONSTERS int16 = 1 << iota
	LD_TWO_SIDED
	LD_UPPER_UNPEGGED
	LD_LOWER_UNPEGGED
	LD_SECRET
	LD_BLOCK_SOUND
	LD_NOT_ON_MAP
	LD_ALREADY_ON_MAP
)

type Linedef struct {
	StartVertexID  int16
	EndVertexID    int16
	Flags          int16
	Type           int16
	TagID          int16
	RightSidedefID int16
	LeftSidedefID  int16
}

func NewLinedefFromBytes(buf []byte) (Linedef, error) {
	if len(buf) != 14 {
		return Linedef{}, fmt.Errorf("invalid linedef entry size; expected 14, received %d", len(buf))
	}
	var ld Linedef
	_, err := binary.Decode(buf, binary.LittleEndian, &ld)
	if err != nil {
		return Linedef{}, fmt.Errorf("error decoding linedef: %v", err)
	}
	return ld, nil
}
