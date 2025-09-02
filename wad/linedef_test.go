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
