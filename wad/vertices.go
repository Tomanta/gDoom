package wad

import (
	"encoding/binary"
	"fmt"
)

type Vertex struct {
	X int16
	Y int16
}

func readVertexFromBuffer(buf []byte) (Vertex, error) {
	var v Vertex
	_, err := binary.Decode(buf, binary.LittleEndian, &v)
	if err != nil {
		return Vertex{}, fmt.Errorf("error decoding vertex: %v", err)
	}

	return v, nil
}

func NewVerticesFromBytes(buf []byte, numEntries int32) ([]Vertex, error) {
	totalVertexSize := numEntries * LUMP_SIZE_VERTEX
	if (int32)(len(buf)) != totalVertexSize {
		return nil, fmt.Errorf("invalid buffer length; expected %d, got %d", totalVertexSize, len(buf))
	}
	var vertices []Vertex

	for i := range numEntries {
		start := i * LUMP_SIZE_VERTEX
		end := start + LUMP_SIZE_VERTEX
		vertex, err := readVertexFromBuffer(buf[start:end])
		if err != nil {
			return nil, fmt.Errorf("error creating vertices: %v", err)
		}
		vertices = append(vertices, vertex)

	}
	return vertices, nil

}
