package main

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"strconv"
	"strings"

	_ "golang.org/x/image/bmp" // Import for BMP decoder

	sa "github.com/CrazyHat62/spriteatlas"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameSprite struct {
	name   string
	pos    rl.Vector2
	region sa.Region
	rect   rl.Rectangle
	frame  int
	step   bool
	loop   bool
}

func main() {
	var ScreenWidth int32 = 1729
	var ScreenHeight int32 = 874
	rl.InitWindow(ScreenWidth, ScreenHeight, "raylib [textures] example - sprite animation")
	defer rl.CloseWindow()

	page, err := sa.Spriteatlas("", "atiles.atlas")
	if err != nil {
		os.Exit(1)
	}
	var img *rl.Image
	targetColor := colorFromStr(page.Alpha_color)
	if page.Alpha_color != "" {
		img = makeImgAlphaTransparent(page.Name, targetColor)
	}
	spriteSheet1 := rl.LoadTextureFromImage(img)
	defer rl.UnloadTexture(spriteSheet1)

	//TODO: Framerate needs fixing
	var gameSpeed int = 1
	FPS := gameSpeed * 4 //4 frames for each tile
	rl.SetTargetFPS(int32(FPS))

	player := GameSprite{name: "player"}
	player.region = page.Regions["player"]
	player.frame = 0
	player.pos.X = 350.0
	player.pos.Y = 100.0

	slime := GameSprite{name: "slime"}
	slime.region = page.Regions["slime_ew"]
	slime.frame = 0
	slime.pos.X = 350.0
	slime.pos.Y = 200.0

	water := GameSprite{name: "water"}
	water.region = page.Regions["region1"]
	water.frame = 0
	water.pos.X = 350.0
	water.pos.Y = 300.0

	explode := GameSprite{name: "explode"}
	explode.region = page.Regions["region5"]
	explode.frame = 0
	explode.pos.X = 350.0 - 48
	explode.pos.Y = 500.0

	for !rl.WindowShouldClose() {

		var rect sa.RECT
		rect, player.frame, player.step, player.loop, err = player.region.GetFrameRect("walk_north", player.frame)
		if err == nil {
			player.rect = rl.Rectangle{X: float32(rect.X), Y: float32(rect.Y), Width: float32(rect.Width), Height: float32(rect.Height)}
		}
		rect, slime.frame, slime.step, slime.loop, err = slime.region.GetFrameRect("east", slime.frame)
		if err == nil {
			slime.rect = rl.Rectangle{X: float32(rect.X), Y: float32(rect.Y), Width: float32(rect.Width), Height: float32(rect.Height)}
		}

		rect, water.frame, water.step, water.loop, err = water.region.GetFrameRect("water", water.frame)
		if err == nil {
			water.rect = rl.Rectangle{X: float32(rect.X), Y: float32(rect.Y), Width: float32(rect.Width), Height: float32(rect.Height)}
		}

		rect, explode.frame, explode.step, explode.loop, err = explode.region.GetFrameRect("explode", explode.frame)
		if err == nil {
			explode.rect = rl.Rectangle{X: float32(rect.X), Y: float32(rect.Y), Width: float32(rect.Width), Height: float32(rect.Height)}
		}

		strw := fmt.Sprintf("%v", water.frame)
		strp := fmt.Sprintf("%v", player.frame)
		strs := fmt.Sprintf("%v", slime.frame)
		stre := fmt.Sprintf("%v", explode.frame)

		rl.BeginDrawing()

		//Background
		rl.ClearBackground(rl.RayWhite)

		rl.DrawTextureRec(spriteSheet1, player.rect, player.pos, rl.White)
		rl.DrawTextureRec(spriteSheet1, water.rect, water.pos, rl.White)
		rl.DrawTextureRec(spriteSheet1, slime.rect, slime.pos, rl.White)
		rl.DrawTextureRec(spriteSheet1, explode.rect, explode.pos, rl.White)

		rl.DrawText(strp, 500.0, 100.0, 40, rl.Black)
		rl.DrawText(strw, 500.0, 200.0, 40, rl.Black)
		rl.DrawText(strs, 500.0, 300.0, 40, rl.Black)
		rl.DrawText(stre, 500.0, 500.0, 40, rl.Black)
		rl.DrawFPS(550, 100)

		rl.EndDrawing()

	}

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
