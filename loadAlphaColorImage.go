package main

import (
	"image"
	"image/color"
	"os"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
	_ "golang.org/x/image/bmp" // Import for BMP decoder
)

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
