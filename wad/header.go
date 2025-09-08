package wad

import (
	"fmt"
)

type Header struct {
	WadType      string
	NumLumps     int32
	DirectoryPos int32
}

// NewHeaderFromBytes takes a 12-byte slice and returns a Header struct
// If the slice is not at least 12 bytes or if it is unable to read
// a piece of the data it will return an error.
func NewHeaderFromBytes(data []byte) (Header, error) {
	if len(data) < 12 {
		return Header{}, fmt.Errorf("file too small: %d", len(data))
	}
	wadType := StringFromBytes(data[0:4])

	numLumps, err := Int32FromBytes(data[4:8])
	if err != nil {
		return Header{}, fmt.Errorf("could not read lump count: %v", err)
	}

	dirPos, err := Int32FromBytes(data[8:12])
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
