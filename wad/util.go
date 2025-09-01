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
