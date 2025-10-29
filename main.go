package main

import (
	"os"

	sa "github.com/CrazyHat62/spriteatlas"
)

func main() {

	page, region, err := sa.Spriteatlas("", "atiles.atlas")
	if err != nil {
		os.Exit(1)
	}
	println(page.PageToStr())
	println(region.RegionToStr())
	println(err)
	for _, key := range region.AnimKeys() {
		println(key)
	}
}
