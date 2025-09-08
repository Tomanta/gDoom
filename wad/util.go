package wad

import (
	"encoding/binary"
	"fmt"
)

// Int16FromBytes takes a slice of 2 bytes in Little Endian encoding and return
// an Int16. LittleEndian is the expected encoding for Doom source files.
// If the provided slice is not exactly 4 bytes it will return an error.
func Int16FromBytes(buf []byte) (int16, error) {
	if len(buf) != 2 {
		return 0, fmt.Errorf("expected 2 bytes, received %d", len(buf))
	}
	return int16(binary.LittleEndian.Uint16(buf)), nil
}

// Int32FromBytes takes a slice of 4 bytes in Little Endian encoding and return
// an Int32. LittleEndian is the expected encoding for Doom source files.
// If the provided slice is not exactly 4 bytes it will return an error.
func Int32FromBytes(buf []byte) (int32, error) {
	if len(buf) != 4 {
		return 0, fmt.Errorf("expected 4 bytes, received %d", len(buf))
	}
	return int32(binary.LittleEndian.Uint32(buf)), nil
}

// StringFromBytes takes a slice of bytes and an expected max length, it will
// return a string from that slice. It will stop parsing as soon as it hits
// a null byte (0x00). If the input buffer is not the expected length it will
// return an error.
func StringFromBytes(buf []byte) string {
	var newBuf []byte
	for _, b := range buf {
		// If we receive a null byte, return the string
		if b == 0x00 {
			return string(newBuf)
		}
		newBuf = append(newBuf, b)
	}
	return string(newBuf)
}
