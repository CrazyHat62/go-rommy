package main

import (
	"fmt"
	"io"
	"os"

	sa "spriteatlas"
)

func main() {
	fmt.Println("Hello World")
	file, reader, err := sa.OpenAtlas("atiles.atlas")
	if err != nil && err != io.EOF {
		os.Exit(1)
	}
	defer file.Close()
	var line string
	line, err = sa.ParseAtlas(reader)
	println(line, err)

}
