package wad

import (
	"encoding/binary"
	"fmt"
)

func Int32FromBytes(buf []byte) (int32, error) {
	if len(buf) != 4 {
		return 0, fmt.Errorf("expected 4 bytes, received %d", len(buf))
	}
	return int32(binary.LittleEndian.Uint32(buf)), nil
}

func StringFromBytes(buf []byte, length int) (string, error) {
	if len(buf) != length {
		return "", fmt.Errorf("expected buffer size %d, received %d", length, len(buf))
	}
	var newBuf []byte
	for _, b := range buf {
		// If we receive a null byte, return the string
		if b == 0x00 {
			return string(newBuf), nil
		}
		newBuf = append(newBuf, b)
	}
	return string(newBuf), nil
}
