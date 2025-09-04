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
	var vertexSize int32 = 4
	if (int32)(len(buf)) != numEntries*vertexSize {
		return nil, fmt.Errorf("invalid buffer length; expected %d, got %d", numEntries*vertexSize, len(buf))
	}
	var vertices []Vertex

	for i := range numEntries {
		start := i * vertexSize
		end := start + vertexSize
		vertex, err := readVertexFromBuffer(buf[start:end])
		if err != nil {
			return nil, fmt.Errorf("error creating vertices: %v", err)
		}
		vertices = append(vertices, vertex)

	}
	return vertices, nil

}
