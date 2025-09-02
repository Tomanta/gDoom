package wad

import (
	"reflect"
	"testing"
)

func TestNewVertexFromBytes(t *testing.T) {
	t.Run("returns error if buffer wrong length", func(t *testing.T) {
		data := []byte{
			0x0c,
		}
		_, err := NewVertexFromBytes(data)
		if err == nil {
			t.Fatalf("did not receive expected error")
		}

	})
	t.Run("returns correct information", func(t *testing.T) {
		data := []byte{
			0x04, 0x17, // 5892
			0xb4, 0xb1, // 	-20044
		}
		got, err := NewVertexFromBytes(data)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		var wantX int16 = 5892
		var wantY int16 = -20044

		if got.X != wantX {
			t.Errorf("want offset %d, got %d", wantX, got.X)
		}
		if got.Y != wantY {
			t.Errorf("want size %d, got %d", wantY, got.Y)
		}
	})
}

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
		for i := range numVertices {
			if !reflect.DeepEqual(got.Vertices[i], want[i]) {
				t.Errorf("wanted %v, got %v", want[i], got.Vertices[i])
			}
		}

	})
}
