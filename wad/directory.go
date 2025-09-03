package wad

import "fmt"

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

type Directory struct {
	Entries []DirEntry
}

// NewDirectoryFromBytes takes a buffer containing the entire directory and a number of
// expected entries. It will return a Directory containing each entry in order.
// If the buffer is not the expect size it will return an error.
func NewDirectoryFromBytes(buf []byte, numEntries int32) (Directory, error) {
	var entrySize int32 = 16
	if (int32)(len(buf)) != numEntries*entrySize {
		return Directory{}, fmt.Errorf("expected buffer size %d (%d * 16 bytes), actual size %d", numEntries*entrySize, numEntries, len(buf))
	}

	var entries []DirEntry

	for i := range numEntries {
		start := i * entrySize
		end := start + entrySize
		entry, err := NewDirEntryFromBytes(buf[start:end])
		if err != nil {
			return Directory{}, fmt.Errorf("error creating directory: %v", err)
		}
		entries = append(entries, entry)

	}
	return Directory{Entries: entries}, nil
}
