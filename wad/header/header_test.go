package header

import (
	"testing"
)

func TestCanLoadHeader(t *testing.T) {
	t.Run("error if buffer too small", func(t *testing.T) {
		testHeader := []byte{
			0x49, 0x57, 0x41, 0x44,
			0xf0, 0x04, 0x00, 0x00,
			0xb4, 0xb7, 0x4f}
		_, err := NewHeaderFromBytes(testHeader)
		if err == nil {
			t.Fatalf("did not receive error")
		}
	})
	t.Run("doom1 shareware header loads correctly", func(t *testing.T) {
		testHeader := []byte{
			0x49, 0x57, 0x41, 0x44, // "IWAD"
			0xf0, 0x04, 0x00, 0x00, // number of lumps: 1264
			0xb4, 0xb7, 0x4f, 0x00, // directory position: 5224372
		}
		got, err := NewHeaderFromBytes(testHeader)
		if err != nil {
			t.Fatalf("Could not read file: %v", err)
		}
		wantType := "IWAD"
		if got.WadType != wantType {
			t.Errorf("wanted header %s, got %s", wantType, got.WadType)
		}

		var wantNumLumps int32 = 1264
		if got.NumLumps != wantNumLumps {
			t.Errorf("wanted number lumps %d, got %d", wantNumLumps, got.NumLumps)
		}

		var wantDirPos int32 = 5224372
		if got.DirectoryPos != wantDirPos {
			t.Errorf("wanted number lumps %d, got %d", wantDirPos, got.DirectoryPos)
		}

	})

	t.Run("pwad header loads correctly", func(t *testing.T) {
		testHeader := []byte{
			0x50, 0x57, 0x41, 0x44, // "PWAD"
			0x02, 0x01, 0x00, 0x00, // number of lumps: 258
			0x04, 0x27, 0x40, 0x00, // directory position: 4204292
		}
		got, err := NewHeaderFromBytes(testHeader)
		if err != nil {
			t.Fatalf("Could not read file: %v", err)
		}
		wantType := "PWAD"
		if got.WadType != wantType {
			t.Errorf("wanted header %s, got %s", wantType, got.WadType)
		}
		var wantNumLumps int32 = 258
		if got.NumLumps != wantNumLumps {
			t.Errorf("wanted number lumps %d, got %d", wantNumLumps, got.NumLumps)
		}

		var wantDirPos int32 = 4204292
		if got.DirectoryPos != wantDirPos {
			t.Errorf("wanted number lumps %d, got %d", wantDirPos, got.DirectoryPos)
		}

	})
}
