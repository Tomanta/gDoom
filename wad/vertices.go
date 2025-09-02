package wad

import "fmt"

type Vertices struct {
	Vertices []Vertex
}

type Vertex struct {
	X int16
	Y int16
}

func NewVertexFromBytes(buf []byte) (Vertex, error) {
	if len(buf) != 4 {
		return Vertex{}, fmt.Errorf("buffer wrong length; expected 4 bytes, received %d", len(buf))
	}
	x, err := Int16FromBytes(buf[0:2])
	if err != nil {
		return Vertex{}, fmt.Errorf("could not read vertex: %v", err)
	}
	y, err := Int16FromBytes(buf[2:4])
	if err != nil {
		return Vertex{}, fmt.Errorf("could not read vertex: %v", err)
	}
	v := Vertex{X: x, Y: y}
	return v, nil
}

func NewVerticesFromBytes(buf []byte, numEntries int32) (Vertices, error) {
	var vertexSize int32 = 4
	if (int32)(len(buf)) != numEntries*vertexSize {
		return Vertices{}, fmt.Errorf("invalid buffer length; expected %d, got %d", numEntries*vertexSize, len(buf))
	}
	var vertices []Vertex

	for i := range numEntries {
		start := i * vertexSize
		end := start + vertexSize
		vertex, err := NewVertexFromBytes(buf[start:end])
		if err != nil {
			return Vertices{}, fmt.Errorf("error creating vertices: %v", err)
		}
		vertices = append(vertices, vertex)

	}
	return Vertices{Vertices: vertices}, nil

}
