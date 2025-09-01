package directory

import (
	"fmt"

	"github.com/tomanta/gdoom/wad"
)

type DirEntry struct {
	Offset int32
	Size   int32
	Name   string
}

func NewDirEntryFromBytes(buf []byte) (DirEntry, error) {
	if len(buf) != 16 {
		return DirEntry{}, fmt.Errorf("invalid directory entry size; expected 16, received %d", len(buf))
	}
	offset, err := wad.Int32FromBytes(buf[0:4])
	if err != nil {
		return DirEntry{}, fmt.Errorf("error reading buffer: %v", err)
	}
	size, err := wad.Int32FromBytes(buf[4:8])
	if err != nil {
		return DirEntry{}, fmt.Errorf("error reading buffer: %v", err)
	}
	name, err := wad.StringFromBytes(buf[8:16], 8)
	if err != nil {
		return DirEntry{}, fmt.Errorf("error reading buffer: %v", err)
	}
	de := DirEntry{
		Offset: offset,
		Size:   size,
		Name:   name,
	}
	return de, nil
}
