package wad

import (
	"reflect"
	"testing"
)

func TestNewLinedefFromBytes(t *testing.T) {
	t.Run("returns error if buffer wrong length", func(t *testing.T) {
		data := []byte{
			0x00, 0x00,
			0x01, 0x00,
			0x01, 0x00,
			0x00, 0x00,
			0x00, 0x00,
			0x00, 0x00,
		}
		_, err := NewLinedefFromBytes(data)
		if err == nil {
			t.Fatalf("did not receive expected error")
		}

	})
	t.Run("returns correct information", func(t *testing.T) {
		data := []byte{
			0x00, 0x00, // Start Vertex: 0
			0x01, 0x00, // End Vertex: 1
			0x01, 0x00, // Flags: 1
			0x00, 0x00, // Type: 0
			0x00, 0x00, // Tag / Trigger: 0
			0x00, 0x00, // Right Sidedef ID: 0. Right is always required, based on start->end.
			0xff, 0xff, // Left sidedef id: none (-1)
		}
		got, err := NewLinedefFromBytes(data)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := Linedef{
			StartVertexID:  0,
			EndVertexID:    1,
			Flags:          1,
			Type:           0,
			TagID:          0,
			RightSidedefID: 0,
			LeftSidedefID:  -1,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}

func TestLinedefHasFlag(t *testing.T) {
	cases := []struct {
		name  string
		data  []byte
		flags int16
		want  bool
	}{
		{"returns true for impassable if 0", []byte{0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF}, LD_IMPASSABLE, true},
		{"returns false for impassable if 1", []byte{0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF}, LD_IMPASSABLE, false},
		{"returns true for block_monsters if 1", []byte{0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF}, LD_BLOCK_MONSTERS, true},
		{"returns false for block_monsters if 4", []byte{0x00, 0x00, 0x01, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF}, LD_BLOCK_MONSTERS, false},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ld, err := NewLinedefFromBytes(tt.data)
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
