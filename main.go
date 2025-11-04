package main

import (
	"image"
	"image/color"
	"os"
	"strconv"
	"strings"

	_ "golang.org/x/image/bmp" // Import for BMP decoder

	sa "github.com/CrazyHat62/spriteatlas"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var ScreenWidth int32 = 1729
	var ScreenHeight int32 = 874
	rl.InitWindow(ScreenWidth, ScreenHeight, "raylib [textures] example - sprite animation")
	defer rl.CloseWindow()

	page, region, err := sa.Spriteatlas("", "atiles.atlas")
	if err != nil {
		os.Exit(1)
	}
	var img *rl.Image
	targetColor := colorFromStr(page.Alpha_color)
	if page.Alpha_color != "" {
		img = makeImgAlphaTransparent(page.Name, targetColor)
	}
	scarfy := rl.LoadTextureFromImage(img)
	defer rl.UnloadTexture(scarfy)

	FPS := 60
	rl.SetTargetFPS(int32(FPS))

	var nextFrame int = 0
	var framesCounter int = 0
	var framesSpeed int = 5
	var rect sa.RECT

	idx := 0
	for !rl.WindowShouldClose() {
		framesCounter++

		if framesCounter >= (60 / framesSpeed) {
			framesCounter = 0

			rect, nextFrame, err = region.GetAnimation("north", nextFrame)
			if err != nil {
				println(err.Error())
				os.Exit(1)
			}

		}

		//rect, err := region.GetAnimation("north", idx)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
		idx++
		frameRec := rl.Rectangle{X: float32(rect.X), Y: float32(rect.Y), Width: float32(rect.Width), Height: float32(rect.Height)}
		rl.BeginDrawing()

		//Background
		rl.ClearBackground(rl.RayWhite)
		rl.DrawTexture(scarfy, 15, 40, rl.RayWhite)
		//lines around the spritesheet
		rl.DrawRectangleLines(15, 40, scarfy.Width, scarfy.Height, rl.Lime)

		// Slide colored squares indicator for speed control
		for i := 0; i < 60; i++ {
			if int32(i) < 60 { //?
				rl.DrawRectangle(int32(250+21*i), 205, 20, 20, rl.Red)
			}
			rl.DrawRectangleLines(int32(250+21*i), 205, 20, 20, rl.Maroon)
		}
		// Show the rectange selection on the sprite
		rl.DrawRectangleLines(15+int32(frameRec.X), 40+int32(frameRec.Y), int32(frameRec.Width), int32(frameRec.Height), rl.Red)
		//position := rl.Vector2{X: 350.0, Y: 280.0}
		// Draw the sprite texture
		rl.DrawTextureRec(scarfy, frameRec, rl.Vector2{X: 350.0, Y: 280.0}, rl.White)

		rl.EndDrawing()
	}

	// println(page.PageToStr())
	// println(region.RegionToStr())
	// if err != nil {
	// 	println("err is " + err.Error())
	// } else {
	// 	println("no errors ")
	// }
	// for _, key := range region.AnimKeys() {
	// 	println(key)
	// }
	// rect1, err := region.GetAnimation("north", 0)
	// rect2, err := region.GetAnimation("north", 1)
	// rect3, err := region.GetAnimation("north", 2)
	// rect4, err := region.GetAnimation("north", 3)

	// rect31, err := region.GetAnimation("west", 0)
	// rect32, err := region.GetAnimation("west", 1)
	// rect33, err := region.GetAnimation("west", 2)
	// rect34, err := region.GetAnimation("west", 3)

	// rect21, err := region.GetAnimation("south", 0)
	// rect22, err := region.GetAnimation("south", 1)
	// rect23, err := region.GetAnimation("south", 2)
	// rect24, err := region.GetAnimation("south", 3)

	// rect11, err := region.GetAnimation("east", 0)
	// rect12, err := region.GetAnimation("east", 1)
	// rect13, err := region.GetAnimation("east", 2)
	// rect14, err := region.GetAnimation("east", 3)

	// if err != nil {
	// 	println(err.Error())
	// } else {
	// 	println(rect1.RectToStr())
	// 	println(rect2.RectToStr())
	// 	println(rect3.RectToStr())
	// 	println(rect4.RectToStr())
	// 	println("")

	// 	println(rect11.RectToStr())
	// 	println(rect12.RectToStr())
	// 	println(rect13.RectToStr())
	// 	println(rect14.RectToStr())
	// 	println("")

	// 	println(rect21.RectToStr())
	// 	println(rect22.RectToStr())
	// 	println(rect23.RectToStr())
	// 	println(rect24.RectToStr())
	// 	println("")

	// 	println(rect31.RectToStr())
	// 	println(rect32.RectToStr())
	// 	println(rect33.RectToStr())
	// 	println(rect34.RectToStr())

	// }
}

func colorFromStr(aColor string) color.RGBA {
	i := strings.Split(aColor, ",")
	r, _ := strconv.Atoi(i[0])
	g, _ := strconv.Atoi(i[1])
	b, _ := strconv.Atoi(i[2])
	a, _ := strconv.Atoi(i[3])
	// 3. Define the color to make transparent (e.g., white)
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)} // White
}

func makeImgAlphaTransparent(filename string, targetColor color.RGBA) *rl.Image {

	file, err := os.Open(filename) // Replace with your image file
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	// Define the target color to make transparent (e.g., white)
	//targetColor = color.RGBA{R: 255, G: 255, B: 255, A: 255} // White

	// Create a new RGBA image to store the result
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	// Iterate through each pixel
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y)
			r, g, b, a := originalColor.RGBA()

			// Convert to 8-bit RGBA for comparison
			r8, g8, b8, a8 := uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8)

			// Check if the color matches the target color
			if r8 == targetColor.R && g8 == targetColor.G && b8 == targetColor.B {
				// Set alpha to 0 for transparency
				newImg.SetRGBA(x, y, color.RGBA{R: r8, G: g8, B: b8, A: 0})
			} else {
				// Keep the original color and alpha
				newImg.SetRGBA(x, y, color.RGBA{R: r8, G: g8, B: b8, A: a8})
			}
		}
	}
	raylibImage := convertImageRGBAtoRaylibImage(newImg)
	return raylibImage
	// // Save the new image with transparency
	// outputFile, err := os.Create("output_transparent.png")
	// if err != nil {
	// 	panic(err)
	// }
	// defer outputFile.Close()

	// png.Encode(outputFile, newImg)
}

func convertImageRGBAtoRaylibImage(goImage *image.RGBA) *rl.Image {
	// Get image dimensions
	width := goImage.Bounds().Dx()
	height := goImage.Bounds().Dy()

	// Create a new raylib.Image
	rlImage := rl.GenImageColor(width, height, rl.Red) // Initialize with a default color

	// Allocate memory for pixel data if not already handled by NewImage
	// Depending on your raylib-go binding, the .Data field might need explicit allocation or be handled by the NewImage function.
	// For example, if .Data is a []byte, you might need:
	// rlImage.Data = make([]byte, width*height*4) // 4 bytes per pixel (RGBA)

	// Iterate through pixels and copy data
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c := goImage.RGBAAt(x, y) // Get the RGBA color at (x, y)

			// Convert Go's color.RGBA to Raylib's Color
			rlColor := rl.NewColor(c.R, c.G, c.B, c.A)

			// Set the pixel in the Raylib image
			rl.ImageDrawPixel(rlImage, int32(x), int32(y), rlColor)
		}
	}

	return rlImage
}
