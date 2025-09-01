package header

import (
	"fmt"

	"github.com/tomanta/gDoom/wad"
	"github.com/tomanta/gdoom/wad"
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
	numLumps, ok := wad.Int32FromBytes(data[4:8]); ok {
		return Header{}, fmt.Errorf("could not read lump count: %v", ok)
	}
	dirPos, _ := wad.Int32FromBytes(data[8:12])
	return Header{Header: header, NumLumps: numLumps, DirectoryPos: dirPos}, nil
}
