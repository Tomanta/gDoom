package wad

import (
	"encoding/binary"
	"fmt"
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
