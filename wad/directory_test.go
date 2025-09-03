package wad

import (
	"reflect"
	"testing"
)

func TestNewDirectoryEntryFromBytes(t *testing.T) {
	t.Run("returns error if buffer wrong length", func(t *testing.T) {
		dirEntry := []byte{
			0x0c, 0x00, 0x00, 0x00, // offset: 12
			0x00, 0x2a, 0x00, 0x00, // lump size: 10752
			0x50, 0x4c, 0x41, 0x59, // name: PLAY
		}
		_, err := NewDirEntryFromBytes(dirEntry)
		if err == nil {
			t.Fatalf("did not receive expected error")
		}

	})
	t.Run("returns correct information", func(t *testing.T) {
		dirEntry := []byte{
			0x0c, 0x00, 0x00, 0x00, // offset: 12
			0x00, 0x2a, 0x00, 0x00, // lump size: 10752
			0x50, 0x4c, 0x41, 0x59, 0x50, 0x41, 0x4c, 0x00, // name: PLAYPAL
		}
		got, err := NewDirEntryFromBytes(dirEntry)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		var wantOffset int32 = 12
		var wantSize int32 = 10752
		wantName := "PLAYPAL"

		if got.Offset != wantOffset {
			t.Errorf("want offset %d, got %d", wantOffset, got.Offset)
		}
		if got.Size != wantSize {
			t.Errorf("want size %d, got %d", wantSize, got.Size)
		}
		if got.Name != wantName {
			t.Errorf("want name %s, got %s", wantName, got.Name)
		}
	})

}

func TestNewDirectoryFromBytes(t *testing.T) {
	t.Run("returns error if buffer length does not match number of entries", func(t *testing.T) {
		entryBuffer := []byte{
			0x0c, 0x00, 0x00, 0x00,
			0x00, 0x2a, 0x00, 0x00,
			0x50, 0x4c, 0x41, 0x59, 0x50, 0x41, 0x4c, 0x00,
			0xff, 0xf2,
		}
		_, err := NewDirectoryFromBytes(entryBuffer, 1)
		if err == nil {
			t.Fatalf("did not receive expected error")
		}
	})
	t.Run("can read a single directory entry", func(t *testing.T) {
		entryBuffer := []byte{
			0x0c, 0x00, 0x00, 0x00, // offset: 12
			0x00, 0x2a, 0x00, 0x00, // lump size: 10752
			0x50, 0x4c, 0x41, 0x59, 0x50, 0x41, 0x4c, 0x00, // name: PLAYPAL
		}
		want := []DirEntry{
			{
				Offset: 12,
				Size:   10752,
				Name:   "PLAYPAL",
			},
		}
		got, _ := NewDirectoryFromBytes(entryBuffer, 1)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %v, got %v", want, got)
		}
	})
	t.Run("can read multiple directory entries", func(t *testing.T) {
		entryBuffer := []byte{
			0x0c, 0x00, 0x00, 0x00, // offset: 12
			0x00, 0x2a, 0x00, 0x00, // lump size: 10752
			0x50, 0x4c, 0x41, 0x59, 0x50, 0x41, 0x4c, 0x00, // name: PLAYPAL
			0x0c, 0x2a, 0x00, 0x00, // offset: 10764
			0x00, 0x22, 0x00, 0x00, // lump size: 8704
			0x43, 0x4f, 0x4c, 0x4f, 0x52, 0x4d, 0x41, 0x50, // name: COLORMAP
		}
		want := []DirEntry{
			{
				Offset: 12,
				Size:   10752,
				Name:   "PLAYPAL",
			},
			{
				Offset: 10764,
				Size:   8704,
				Name:   "COLORMAP",
			},
		}
		got, err := NewDirectoryFromBytes(entryBuffer, 2)
		if err != nil {
			t.Fatalf("error reading directory, %v", err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("wanted %v, got %v", want, got)
		}
	})
}
