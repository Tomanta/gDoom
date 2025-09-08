package wad

import (
	"reflect"
	"testing"
)

func TestNewSectorsFromBytes(t *testing.T) {
	t.Run("returns error if buffer wrong length", func(t *testing.T) {
		data := []byte{
			0x00, 0x00,
			0x00, 0x00,
			0x2D, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x2D, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x44, 0x4F,
			0x4F, 0x52,
		}
		var numSectors int32 = 2
		_, err := NewSectorsFromBytes(data, numSectors)
		if err == nil {
			t.Fatalf("did not receive expected error")
		}
	})

	t.Run("returns correct information", func(t *testing.T) {
		data := []byte{
			0x00, 0x00, // Floor height
			0x48, 0x00, // Ceiling height
			0x46, 0x4C, 0x4F, 0x4F, 0x52, 0x34, 0x5F, 0x38, // Floor Texture; FLOOR4_8
			0x43, 0x45, 0x49, 0x4C, 0x33, 0x5F, 0x35, 0x00, // Ceiling Texture; CEIL3_5
			0xA0, 0x00, // Light level
			0x00, 0x00, // Special Sector Type
			0x00, 0x00, // Tag number

			0x20, 0x00,
			0x58, 0x00,
			0x46, 0x4C, 0x41, 0x54, 0x31, 0x38, 0x00, 0x00, // FLAT18
			0x43, 0x45, 0x49, 0x4C, 0x35, 0x5F, 0x31, 0x00, // CEIL5_1
			0xFF, 0x00,
			0x00, 0x00,
			0x00, 0x00,
		}
		var numSectors int32 = 2
		want := []Sector{
			{
				FloorHeight:    0,
				CeilingHeight:  72,
				FloorTexture:   "FLOOR4_8",
				CeilingTexture: "CEIL3_5",
				LightLevel:     160,
				SpecialType:    0,
				TagID:          0,
			}, {
				FloorHeight:    32,
				CeilingHeight:  88,
				FloorTexture:   "FLAT18",
				CeilingTexture: "CEIL5_1",
				LightLevel:     255,
				SpecialType:    0,
				TagID:          0,
			},
		}
		got, err := NewSectorsFromBytes(data, numSectors)
		if err != nil {
			t.Fatalf("could not read sectors: %v", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %v, got %v", want, got)
		}

	})
}
