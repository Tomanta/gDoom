package wad

import (
	"reflect"
	"testing"
)

func TestNewThingsFromBytes(t *testing.T) {
	t.Run("returns error if buffer wrong length", func(t *testing.T) {
		data := []byte{
			0x20, 0x04, // XPos
			0xe0, 0xf1, // YPos
			0x5a, 0x00, // Angle
			0x01, 0x00, // Type
			0x07, 0x00, // Flags
			0xf0, 0x03,
			0xf0, 0xf1,
			0x5a, 0x00,
		}
		var numThings int32 = 2
		_, err := NewThingsFromBytes(data, numThings)
		if err == nil {
			t.Fatalf("did not receive expected error")
		}
	})

	t.Run("returns correct information", func(t *testing.T) {
		data := []byte{
			0x20, 0x04, // XPos
			0xe0, 0xf1, // YPos
			0x5a, 0x00, // Angle
			0x01, 0x00, // Type
			0x07, 0x00, // Flags
			0xf0, 0x03,
			0xf0, 0xf1,
			0x5a, 0x00,
			0x02, 0x00,
			0x07, 0x00}
		var numThings int32 = 2
		want := []Thing{
			{XPos: 1056, YPos: -3616, Angle: 90, TypeID: 1, Flags: 7},
			{XPos: 1008, YPos: -3600, Angle: 90, TypeID: 2, Flags: 7},
		}
		got, err := NewThingsFromBytes(data, numThings)
		if err != nil {
			t.Fatalf("could not read Things: %v", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %v, got %v", want, got)
		}

	})
}
