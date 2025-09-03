package wad

import (
	"reflect"
	"testing"
)

func TestNewVerticesFromBytes(t *testing.T) {
	t.Run("returns error if buffer wrong length", func(t *testing.T) {
		data := []byte{
			0x0c, 0x42, 0xb4, 0x23, 0x45, 0x21,
		}
		var numVertices int32 = 2
		_, err := NewVerticesFromBytes(data, numVertices)
		if err == nil {
			t.Fatalf("did not receive expected error")
		}
	})

	t.Run("returns correct information", func(t *testing.T) {
		data := []byte{
			0x04, 0x17, //  5892
			0xb4, 0xb1, // -20044
			0x04, 0x29, //  10500
			0xb1, 0x11, //	4529
		}
		var numVertices int32 = 2
		want := []Vertex{
			{X: 5892, Y: -20044},
			{X: 10500, Y: 4529},
		}
		got, err := NewVerticesFromBytes(data, numVertices)
		if err != nil {
			t.Fatalf("could not read vertices: %v", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %v, got %v", want, got)
		}

	})
}
