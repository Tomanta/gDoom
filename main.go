package main

import (
	"fmt"
	"os"

	"github.com/tomanta/gdoom/wad"
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

}
