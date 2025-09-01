package wad

import "testing"

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
