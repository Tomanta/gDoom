package wad

import "fmt"

type DirEntry struct {
	Offset int32
	Size   int32
	Name   string
}

func readDirEntryFromBuffer(buf []byte) (DirEntry, error) {
	offset, err := Int32FromBytes(buf[0:4])
	if err != nil {
		return DirEntry{}, fmt.Errorf("error reading buffer for lump offset: %v", err)
	}
	size, err := Int32FromBytes(buf[4:8])
	if err != nil {
		return DirEntry{}, fmt.Errorf("error reading buffer for lump size: %v", err)
	}
	name := StringFromBytes(buf[8:16])

	de := DirEntry{
		Offset: offset,
		Size:   size,
		Name:   name,
	}
	return de, nil
}

// NewDirectoryFromBytes takes a buffer containing the entire directory and a number of
// expected entries. It will return a Directory containing each entry in order.
// If the buffer is not the expect size it will return an error.
func NewDirectoryFromBytes(buf []byte, numEntries int32) ([]DirEntry, error) {
	var entrySize int32 = 16
	if (int32)(len(buf)) != numEntries*entrySize {
		return []DirEntry{}, fmt.Errorf("expected buffer size %d (%d * 16 bytes), actual size %d", numEntries*entrySize, numEntries, len(buf))
	}

	var directory []DirEntry

	for i := range numEntries {
		start := i * entrySize
		end := start + entrySize
		entry, err := readDirEntryFromBuffer(buf[start:end])
		if err != nil {
			return []DirEntry{}, fmt.Errorf("error creating directory: %v", err)
		}
		directory = append(directory, entry)

	}
	return directory, nil
}
