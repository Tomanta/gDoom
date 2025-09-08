package wad

import (
	"slices"
)

type Level struct {
	Name     string
	Vertices []Vertex
	Linedefs []Linedef
	Sidedefs []Sidedef
	Sectors  []Sector
	Things   []Thing
}

func isLevel(name string) bool {
	if len(name) == 4 && name[0] == 'E' && name[2] == 'M' {
		return true
	}
	if len(name) == 5 && name[0:3] == "MAP" {
		return true
	}
	return false
}

func isLevelLump(name string) bool {
	var levelLumps = []string{"THINGS", "LINEDEFS", "SIDEDEFS", "VERTEXES", "SEGS", "SSECTORS", "NODES", "SECTORS", "REJECT", "BLOCKMAP"}

	return slices.Contains(levelLumps, name)
}

func NewLevelFromBuffer(buf []byte, directory []DirEntry) (Level, error) {
	level := Level{}
	for _, de := range directory {
		data := buf[de.Offset : de.Offset+de.Size]
		switch {
		case isLevel(de.Name):
			level.Name = de.Name

		case de.Name == "VERTEXES":
			level.Vertices, _ = NewVerticesFromBytes(data, de.Size/LUMP_SIZE_VERTEX)

		case de.Name == "LINEDEFS":
			level.Linedefs, _ = NewLinedefsFromBytes(data, de.Size/LUMP_SIZE_LINEDEF)
		case de.Name == "SIDEDEFS":
			level.Sidedefs, _ = NewSidedefsFromBytes(data, de.Size/LUMP_SIZE_SIDEDEF)
		case de.Name == "SECTORS":
			level.Sectors, _ = NewSectorsFromBytes(data, de.Size/LUMP_SIZE_SECTOR)
		case de.Name == "THINGS":
			level.Things, _ = NewThingsFromBytes(data, de.Size/LUMP_SIZE_THING)
		}
	}
	return level, nil
}

// GetVertexOffsets returns the adjustments needed for each vertex x and y position
// to convert the DOOM map data (which can have negative positions) into something that
// can be drawn on screen, such as for an automap.
func (l Level) GetVertexOffsets() (int32, int32) {
	var x int16 = 0
	var y int16 = 0

	for _, v := range l.Vertices {
		if v.X < x {
			x = v.X
		}
		if v.Y < y {
			y = v.Y
		}
	}
	return (int32)(x * -1), (int32)(y * -1)
}
