package wad

import (
	"reflect"
	"testing"
)

func TestLinedefHasFlag(t *testing.T) {
	cases := []struct {
		name  string
		data  []byte
		flags int16
		want  bool
	}{
		{"returns true for impassable if 1", []byte{0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF}, ML_BLOCKING, true},
		{"returns false for impassable if 4", []byte{0x00, 0x00, 0x01, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF}, ML_BLOCKING, false},
		{"returns true for impassable if 3", []byte{0x00, 0x00, 0x01, 0x00, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF}, ML_BLOCKING, true},
		{"returns true for block_monsters if 2", []byte{0x00, 0x00, 0x02, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF}, ML_BLOCKMONSTERS, true},
		{"returns false for block_monsters if 4", []byte{0x00, 0x00, 0x01, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF}, ML_BLOCKMONSTERS, false},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ld, err := readLinedefFromBuffer(tt.data)
			if err != nil {
				t.Fatalf("could not create linedef: %v", err)
			}
			got := ld.HasFlag(tt.flags)
			if got != tt.want {
				t.Errorf("want %t, got %t", tt.want, got)
			}
		})
	}
}

func TestNewLinedefsFromBytes(t *testing.T) {
	t.Run("returns error if buffer wrong length", func(t *testing.T) {
		data := []byte{
			0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF,
			0x01, 0x00, 0x02, 0x00, 0x01, 0x00, 0x00,
		}
		var numLinedefs int32 = 2
		_, err := NewLinedefsFromBytes(data, numLinedefs)
		if err == nil {
			t.Fatalf("did not receive expected error")
		}
	})

	t.Run("returns correct information", func(t *testing.T) {
		data := []byte{
			0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF,
			0x01, 0x00, 0x02, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0xFF, 0xFF,
		}
		var numLinedefs int32 = 2
		want := []Linedef{
			{
				StartVertexID:  0,
				EndVertexID:    1,
				Flags:          1,
				Type:           0,
				TagID:          0,
				RightSidedefID: 0,
				LeftSidedefID:  -1,
			}, {
				StartVertexID:  1,
				EndVertexID:    2,
				Flags:          1,
				Type:           0,
				TagID:          0,
				RightSidedefID: 1,
				LeftSidedefID:  -1,
			},
		}
		got, err := NewLinedefsFromBytes(data, numLinedefs)
		if err != nil {
			t.Fatalf("could not read linedefs: %v", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %v, got %v", want, got)
		}

	})
}
