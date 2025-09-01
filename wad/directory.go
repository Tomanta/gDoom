package wad

import "fmt"

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
