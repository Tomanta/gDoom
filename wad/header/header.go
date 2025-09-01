package header

import (
	"fmt"

	"github.com/tomanta/gdoom/wad"
)

type Header struct {
	WadType      string
	NumLumps     int32
	DirectoryPos int32
}

func NewHeaderFromBytes(data []byte) (Header, error) {
	if len(data) < 12 {
		return Header{}, fmt.Errorf("file too small: %d", len(data))
	}
	wadType := string(data[0:4])

	numLumps, err := wad.Int32FromBytes(data[4:8])
	if err != nil {
		return Header{}, fmt.Errorf("could not read lump count: %v", err)
	}

	dirPos, err := wad.Int32FromBytes(data[8:12])
	if err != nil {
		return Header{}, fmt.Errorf("could not read directory position: %v", err)
	}

	header := Header{
		WadType:      wadType,
		NumLumps:     numLumps,
		DirectoryPos: dirPos,
	}

	return header, nil
}
