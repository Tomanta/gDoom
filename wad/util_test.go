package wad

import (
	"testing"
)

func TestInt32FromBytes(t *testing.T) {
	t.Run("generates error if buffer too short", func(t *testing.T) {
		buf := []byte{0x3f}
		_, err := Int32FromBytes(buf)
		if err == nil {
			t.Errorf("expected invalid buffer size error")
		}
	})
	t.Run("generates error if buffer too long", func(t *testing.T) {
		buf := []byte{0x3f, 0x42, 0x35, 0xff, 0xab}
		_, err := Int32FromBytes(buf)
		if err == nil {
			t.Errorf("expected invalid buffer size error")
		}
	})

	t.Run("returns correct int32", func(t *testing.T) {
		buf := []byte{0xb4, 0xb7, 0x3f, 0x00}
		var want int32 = 4175796
		got, err := Int32FromBytes(buf)
		if err != nil {
			t.Fatalf("received unexpected error %v", err)
		}
		if got != want {
			t.Errorf("wanted %d, got %d", want, got)
		}
	})
}

func TestStringFromBytes(t *testing.T) {
	buf := []byte{0x50, 0x4c, 0x41, 0x59, 0x50, 0x41, 0x4c, 0x00} // PLAYPAL
	got, _ := StringFromBytes(buf, 8)
	wantLen := 7
	want := "PLAYPAL"
	if len(got) != wantLen {
		t.Errorf("wanted string of length %d, got %d", wantLen, len(got))
	}
	if want != got {
		t.Errorf("expected string %s, got %s", want, got)
	}
}
