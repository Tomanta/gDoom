package wad

import (
	"fmt"
)

type DirEntry struct {
	Offset int32
	Size   int32
	Name   string
}

// NewDirEntryFromBytes takes a []byte buffer of 16 characters and returns a new DirEntry.
// if the buffer is not exactly 16 bytes, returns an error.
func NewDirEntryFromBytes(buf []byte) (DirEntry, error) {
	if len(buf) != 16 {
		return DirEntry{}, fmt.Errorf("invalid directory entry size; expected 16, received %d", len(buf))
	}
	offset, err := Int32FromBytes(buf[0:4])
	if err != nil {
		return DirEntry{}, fmt.Errorf("error reading buffer for lump offset: %v", err)
	}
	size, err := Int32FromBytes(buf[4:8])
	if err != nil {
		return DirEntry{}, fmt.Errorf("error reading buffer for lump size: %v", err)
	}
	name, err := StringFromBytes(buf[8:16], 8)
	if err != nil {
		return DirEntry{}, fmt.Errorf("error reading buffer for lump name: %v", err)
	}
	de := DirEntry{
		Offset: offset,
		Size:   size,
		Name:   name,
	}
	return de, nil
}
