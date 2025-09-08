package main

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"github.com/tomanta/gdoom/wad"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
)

func main() {
	testFile := "./gamefiles/doom1.wad"
	data, err := os.ReadFile(testFile)
	if err != nil {
		panic("could not open wad file")
	}

	wad, err := wad.NewWadFromBytes(data)
	if err != nil {
		fmt.Printf("could not create wad: %v", err)
		return
	}

	fmt.Printf("DEBUG: Number of lumps: %d\n", wad.Header.NumLumps)
	fmt.Printf("DEBUG: Number of directory entries: %d\n", len(wad.Directory))
	/*
		fmt.Printf("Searcing for sectors...\n")
		for _, i := range wad.Directory {
			if i.Name == "SECTORS" {
				fmt.Printf("Sectors found!\n")
				fmt.Printf("DEBUG: Sector 1: %s\n", data[i.Offset:i.Offset+26])
				fmt.Printf("DEBUG: Sector 1: %s\n", data[i.Offset+26:i.Offset+(26*2)])
				break
			}
		}
	*/
	drawE1M1(wad.Levels[0])
}

func drawE1M1(l wad.Level) {
	offx, offy := l.GetVertexOffsets()

	type vtx struct {
		X int32
		Y int32
	}
	type line struct {
		start_v vtx
		end_v   vtx
		color   color.Color
	}
	var lines []line

	for _, ld := range l.Linedefs {
		sv := vtx{
			X: (int32)(l.Vertices[ld.StartVertexID].X) + offx,
			Y: (int32)(l.Vertices[ld.StartVertexID].Y) + offy,
		}
		ev := vtx{
			X: (int32)(l.Vertices[ld.EndVertexID].X) + offx,
			Y: (int32)(l.Vertices[ld.EndVertexID].Y) + offy,
		}
		var c color.RGBA
		if ld.LeftSidedefID == -1 {
			c = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Red
		} else if l.Sectors[l.Sidedefs[ld.LeftSidedefID].SectorID].CeilingHeight != l.Sectors[l.Sidedefs[ld.RightSidedefID].SectorID].CeilingHeight {
			c = color.RGBA{R: 255, G: 255, B: 0, A: 255} // Yellow
		} else if l.Sectors[l.Sidedefs[ld.LeftSidedefID].SectorID].FloorHeight != l.Sectors[l.Sidedefs[ld.RightSidedefID].SectorID].FloorHeight {
			c = color.RGBA{R: 165, G: 42, B: 42, A: 255} // Brown

		} else {
			c = color.RGBA{R: 255, G: 255, B: 255, A: 255} // White
		}
		lines = append(lines, line{
			start_v: sv,
			end_v:   ev,
			color:   c,
		})
	}

	var max_x int32 = 0
	var max_y int32 = 0
	for _, l := range lines {
		if l.start_v.X > max_x {
			max_x = l.start_v.X
		}
		if l.start_v.Y > max_y {
			max_y = l.start_v.Y
		}

		if l.end_v.X > max_x {
			max_x = l.end_v.X
		}
		if l.end_v.Y > max_y {
			max_y = l.end_v.Y
		}
	}

	dest := image.NewRGBA(image.Rect(0, 0, (int)(max_x), (int)(max_y)))
	gc := draw2dimg.NewGraphicContext(dest)

	gc.SetFillColor(image.Black)
	draw2dkit.Rectangle(gc, 0, 0, (float64)(max_x), (float64)(max_x))
	gc.Fill()

	for _, l := range lines {
		DrawLine(gc, l.color, (int)(l.start_v.X), (int)(l.start_v.Y), (int)(l.end_v.X), (int)(l.end_v.Y))
	}
	draw2dimg.SaveToPngFile("map.png", dest)
}

func DrawLine(gc draw2d.GraphicContext, c color.Color, x0, y0, x1, y1 int) {
	gc.SetStrokeColor(c)
	gc.MoveTo((float64)(x0), (float64)(y0))
	gc.LineTo((float64)(x1), (float64)(y1))
	gc.Stroke()
}
