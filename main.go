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
	rect1, err := sa.GetAnimation(region, "north", 0)
	rect2, err := sa.GetAnimation(region, "north", 1)
	rect3, err := sa.GetAnimation(region, "north", 2)
	rect4, err := sa.GetAnimation(region, "north", 3)

	rect31, err := sa.GetAnimation(region, "west", 0)
	rect32, err := sa.GetAnimation(region, "west", 1)
	rect33, err := sa.GetAnimation(region, "west", 2)
	rect34, err := sa.GetAnimation(region, "west", 3)

	rect21, err := sa.GetAnimation(region, "south", 0)
	rect22, err := sa.GetAnimation(region, "south", 1)
	rect23, err := sa.GetAnimation(region, "south", 2)
	rect24, err := sa.GetAnimation(region, "south", 3)

	rect11, err := sa.GetAnimation(region, "east", 0)
	rect12, err := sa.GetAnimation(region, "east", 1)
	rect13, err := sa.GetAnimation(region, "east", 2)
	rect14, err := sa.GetAnimation(region, "east", 3)


	if err != nil {
		println(err.Error())
	} else {
		println(rect1.X1, rect1.Y1, rect1.X2, rect1.Y2)
		println(rect2.X1, rect2.Y1, rect2.X2, rect2.Y2)
		println(rect3.X1, rect3.Y1, rect3.X2, rect3.Y2)
		println(rect4.X1, rect4.Y1, rect4.X2, rect4.Y2)
		println("")

		println(rect11.X1, rect11.Y1, rect11.X2, rect11.Y2)
		println(rect12.X1, rect12.Y1, rect12.X2, rect12.Y2)
		println(rect13.X1, rect13.Y1, rect13.X2, rect13.Y2)
		println(rect14.X1, rect14.Y1, rect14.X2, rect14.Y2)
		println("")

		println(rect21.X1, rect21.Y1, rect21.X2, rect21.Y2)
		println(rect22.X1, rect22.Y1, rect22.X2, rect22.Y2)
		println(rect23.X1, rect23.Y1, rect23.X2, rect23.Y2)
		println(rect24.X1, rect24.Y1, rect24.X2, rect24.Y2)
		println("")

		println(rect31.X1, rect31.Y1, rect31.X2, rect31.Y2)
		println(rect32.X1, rect32.Y1, rect32.X2, rect32.Y2)
		println(rect33.X1, rect33.Y1, rect33.X2, rect33.Y2)
		println(rect34.X1, rect34.Y1, rect34.X2, rect34.Y2)

	}
}
