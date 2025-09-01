package header

import (
	"encoding/binary"
	"fmt"
)

type Header struct {
	Header       string
	NumLumps     int32
	DirectoryPos int32
}

func NewHeaderFromBytes(data []byte) (Header, error) {
	if len(data) < 12 {
		return Header{}, fmt.Errorf("file too small: %d", len(data))
	}
	header := string(data[0:4])
	numLumps := int32(binary.LittleEndian.Uint32(data[4:8]))
	dirPos := int32(binary.LittleEndian.Uint32(data[8:12]))
	return Header{Header: header, NumLumps: numLumps, DirectoryPos: dirPos}, nil
}
