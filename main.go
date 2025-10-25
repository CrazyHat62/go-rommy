package main

import (
	"os"

	sa "spriteatlas"
)

func main() {

	err := sa.Spriteatlas("", "atiles.atlas")
	if err != nil {
		os.Exit(1)
	}

}
