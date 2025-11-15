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
	Name   string
	Pos    rl.Vector2
	Region sa.Region
	Rect   rl.Rectangle
	Frame  int
	step   bool
	Loop   bool
	Played bool
}

var frameCounter int

func (g *GameSprite) Init(name string, region string, X float32, Y float32) {
	g.Name = name
	g.Region = page.Regions[region]
	g.Frame = 0
	g.Pos.X = X
	g.Pos.Y = Y
}
func (g *GameSprite) X() float32 {
	return g.Pos.X
}
func (g *GameSprite) Y() float32 {
	return g.Pos.Y
}

func (g *GameSprite) centerX() float32 {
	return g.X() + g.Width()/2.0
}

func (g *GameSprite) centerY() float32 {
	return g.Y() + g.Height()/2.0
}

func (g *GameSprite) SetX(x float32) {
	g.Pos.X = x
}

func (g *GameSprite) SetY(y float32) {
	g.Pos.Y = y
}

func (g *GameSprite) Width() float32 {
	return g.Rect.Width
}

func (g *GameSprite) Height() float32 {
	return g.Rect.Height
}
func (g *GameSprite) Update(anim string) error {
	var rect sa.RECT
	var err error
	rect, g.Frame, g.step, g.Loop, err = g.Region.GetFrameRect(anim, g.Frame)
	if err == nil {
		g.Rect = rl.Rectangle{X: float32(rect.X), Y: float32(rect.Y), Width: float32(rect.Width), Height: float32(rect.Height)}
	}
	return err
}

func (g *GameSprite) Step(dir int, stepSize int, speed int) {

	if g.step && g.Frame == 0 { //animation does the movement
		switch dir {
		case 0: //north
			g.SetY(g.Y() + float32(stepSize))
		case 1: //east
			g.SetX(g.X() + float32(stepSize))
		case 2: //south
			g.SetY(g.Y() - float32(stepSize))
		case 3: //west
			g.SetX(g.X() - float32(stepSize))
		}
	}
	if !g.step {
		switch dir {
		case 0: //north
			g.SetY(g.Y() - g.Height()/float32(speed))
		case 1: //east
			g.SetX(g.X() + g.Width()/float32(speed))
		case 2: //south
			g.SetY(g.Y() + g.Height()/float32(speed))
		case 3: //west
			g.SetX(g.X() - g.Width()/float32(speed))
		}

	}

}

var WorldWidth int32 = 1536
var WorldHeight int32 = 1536
var ScreenWidth int32 = 1729 //432
var ScreenHeight int32 = 874 //432

var spriteSheet1 rl.Texture2D
var page *sa.Page
var err error

func init() {

	rl.InitWindow(ScreenWidth, ScreenHeight, "raylib [textures] example - sprite animation")
	page, err = sa.Spriteatlas("", "atiles.atlas")
	if err != nil {
		os.Exit(1)
	}
	var img *rl.Image
	targetColor := colorFromStr(page.Alpha_color)
	if page.Alpha_color != "" {
		img = makeImgAlphaTransparent(page.Name, targetColor)
	}
	spriteSheet1 = rl.LoadTextureFromImage(img)
}

func main() {

	defer rl.CloseWindow()
	defer rl.UnloadTexture(spriteSheet1)

	//TODO: Framerate needs fixing
	gameSpeed := 4
	FPS := 4
	rl.SetTargetFPS(int32(FPS))

	var player GameSprite
	var slime GameSprite
	var water GameSprite
	var explode GameSprite
	var tile GameSprite

	player.Init("player", "player", 336.0, 576.0)
	slime.Init("slime", "slime_ew", 336.0, 192.0)
	water.Init("water", "region1", 336.0, 288.0)
	explode.Init("explode", "region5", 336.0-48, 480.0)
	tile.Init("tile", "region1", 0.0, 0.0)
	err = tile.Update("tile")

	target := rl.LoadRenderTexture(WorldWidth, WorldHeight)
	defer rl.UnloadRenderTexture(target)

	//create background texture
	rl.BeginTextureMode(target)

	for x := 0; x < 32; x++ {
		for y := 0; y < 32; y++ {
			tile.SetX(float32(x * 48))
			tile.SetY(float32(y * 48))
			rl.DrawTextureRec(spriteSheet1, tile.Rect, tile.Pos, rl.White)
		}
	}

	rl.EndTextureMode()

	camera := rl.Camera2D{}
	var camTarget *GameSprite = &slime
	camera.Target = rl.NewVector2(camTarget.centerX(), camTarget.centerY())
	camera.Offset = rl.NewVector2(float32(ScreenWidth/2), float32(ScreenHeight/2))
	camera.Rotation = 0.0
	camera.Zoom = 1.0

	for !rl.WindowShouldClose() {

		strw := fmt.Sprintf("%v", water.Frame)

		strs := fmt.Sprintf("%v", slime.Frame)
		stre := fmt.Sprintf("%v", explode.Frame)

		if frameCounter > gameSpeed/FPS {
			frameCounter = 0
			err = player.Update("walk_north")
			err = slime.Update("east")
			err = water.Update("water")
			err = explode.Update("explode")
		} else {
			frameCounter++
		}

		camera.Target = rl.NewVector2(camTarget.centerX(), camTarget.centerY())

		rl.BeginDrawing()

		//Background
		rl.ClearBackground(rl.Black)
		rl.BeginMode2D(camera)

		t := target.Texture
		rl.DrawTextureRec(
			t,
			rl.Rectangle{X: 0, Y: 0, Width: float32(t.Width), Height: -float32(t.Height)}, // Flip vertically
			rl.Vector2{X: 0, Y: 0},
			rl.White,
		)

		rl.DrawTextureRec(spriteSheet1, player.Rect, player.Pos, rl.White)
		rl.DrawTextureRec(spriteSheet1, water.Rect, water.Pos, rl.White)
		rl.DrawTextureRec(spriteSheet1, slime.Rect, slime.Pos, rl.White)
		if !explode.Played {
			rl.DrawTextureRec(spriteSheet1, explode.Rect, explode.Pos, rl.White)
		}
		strp := fmt.Sprintf("%v", int32(player.centerY()))
		rl.DrawText(strp, int32(player.X())-10, int32(player.Y())-20, 20, rl.White)

		rl.EndMode2D()

		rl.DrawText(strw, 500.0, 200.0, 40, rl.Black)
		rl.DrawText(strs, 500.0, 300.0, 40, rl.Black)
		if !explode.Played {
			rl.DrawText(stre, 500.0, 500.0, 40, rl.Black)
		}

		rl.DrawFPS(550, 100)

		rl.EndDrawing()

		//step upwards

		player.Step(0, 48, gameSpeed)

		slime.Step(1, 48, gameSpeed)

		// if player.Step {
		// 	if player.Frame == 0 {
		// 		player.Pos.Y = player.Y() - float32(player.Height()) + 48
		// 	}
		// } else {
		// 	player.Pos.Y -= player.Height() / float32(gameSpeed)
		// }

		// if slime.Step {
		// 	if slime.Frame == 0 {
		// 		slime.setX(slime.X() + slime.Width() - 48)
		// 	}
		// } else {
		// 	slime.setX(slime.X() + slime.Width()/float32(FPS))
		// }

		if explode.Loop == false && explode.Frame == 0 {
			explode.Played = true
		}

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
