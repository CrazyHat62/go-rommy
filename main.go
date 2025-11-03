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
	if err != nil {
		println("err is " + err.Error())
	} else {
		println("no errors ")
	}
	for _, key := range region.AnimKeys() {
		println(key)
	}
	rect1, err := region.GetAnimation("north", 0)
	rect2, err := region.GetAnimation("north", 1)
	rect3, err := region.GetAnimation("north", 2)
	rect4, err := region.GetAnimation("north", 3)

	rect31, err := region.GetAnimation("west", 0)
	rect32, err := region.GetAnimation("west", 1)
	rect33, err := region.GetAnimation("west", 2)
	rect34, err := region.GetAnimation("west", 3)

	rect21, err := region.GetAnimation("south", 0)
	rect22, err := region.GetAnimation("south", 1)
	rect23, err := region.GetAnimation("south", 2)
	rect24, err := region.GetAnimation("south", 3)

	rect11, err := region.GetAnimation("east", 0)
	rect12, err := region.GetAnimation("east", 1)
	rect13, err := region.GetAnimation("east", 2)
	rect14, err := region.GetAnimation("east", 3)

	if err != nil {
		println(err.Error())
	} else {
		println(rect1.RectToStr())
		println(rect2.RectToStr())
		println(rect3.RectToStr())
		println(rect4.RectToStr())
		println("")

		println(rect11.RectToStr())
		println(rect12.RectToStr())
		println(rect13.RectToStr())
		println(rect14.RectToStr())
		println("")

		println(rect21.RectToStr())
		println(rect22.RectToStr())
		println(rect23.RectToStr())
		println(rect24.RectToStr())
		println("")

		println(rect31.RectToStr())
		println(rect32.RectToStr())
		println(rect33.RectToStr())
		println(rect34.RectToStr())

	}
}
