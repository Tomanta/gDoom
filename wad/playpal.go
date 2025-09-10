package wad

import (
	"fmt"
	"image/color"
)

type Palette struct {
	Colors [256]color.Color
}

func NewPlaypalFromBytes(buf []byte) ([14]Palette, error) {
	totalPlaypalSize := LUMP_NUM_PALETTES * LUMP_SIZE_PALETTE
	if len(buf) != LUMP_NUM_PALETTES*LUMP_SIZE_PALETTE {
		return [14]Palette{}, fmt.Errorf("invalid buffer length; expected %d, got %d", totalPlaypalSize, len(buf))
	}
	// 14 palettes
	palettes := [14]Palette{}
	for i := range 14 {
		palettes[i] = readPaletteFromBytes(buf[i*LUMP_SIZE_PALETTE : i*LUMP_SIZE_PALETTE+LUMP_SIZE_PALETTE])
	}
	return palettes, nil
}

func readPaletteFromBytes(buf []byte) Palette {
	// 256 colors per palette. 3 bytes per color, Red, Green, and Blue
	var colors [256]color.Color
	for i := range 256 {
		r := Uint8FromByte(buf[i*3])
		g := Uint8FromByte(buf[i*3+1])
		b := Uint8FromByte(buf[i*3+2])
		colors[i] = color.RGBA{r, g, b, 255}
	}
	return Palette{Colors: colors}
}
