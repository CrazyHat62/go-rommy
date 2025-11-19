package main

import (
	"fmt"
	"os"

	sa "github.com/CrazyHat62/spriteatlas"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameSprite struct {
	Name         string
	Pos          rl.Vector2
	Region       sa.Region
	Rect         rl.Rectangle
	CurrentAnim  sa.Anim
	CurrentFrame int
	Played       bool
	//	timeBetweenFrames float32
}

func (g *GameSprite) Init(name string, region string, X float32, Y float32) {
	g.Name = name
	g.Region = page.Regions[region]
	g.CurrentFrame = 0
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

func (g *GameSprite) UpdateFrame(animName string, fnum int) error {
	var rect sa.RECT

	reg := &g.Region

	anim, err := reg.GetAnimation(animName)
	if err != nil {
		return err
	}
	g.CurrentAnim = anim
	rect, err = reg.GetFrameRect(anim, fnum)
	if err != nil {
		return err
	}
	g.Rect = rl.Rectangle{X: float32(rect.X), Y: float32(rect.Y), Width: float32(rect.Width), Height: float32(rect.Height)}
	return err
}

func (g *GameSprite) StepDistance(dir int, stepSize int, speed int) {

	// if Step the animate once per animation
	if g.CurrentAnim.Step && g.CurrentFrame == 0 {
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
	if !g.CurrentAnim.Step {
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
var ScreenWidth int32 = 1536  //480  //432
var ScreenHeight int32 = 1536 //480 //432

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

	var gameSpeed float32 = 5
	var animSpeed float32 = 4
	// use some standard 1 unit of measurement / frames per animation
	var timeBetweenFrames float32 = animSpeed / gameSpeed
	var accumulatedTime float32 = 0.0

	rl.SetTargetFPS(60)

	var player GameSprite
	var slime GameSprite
	var water GameSprite
	var explode GameSprite
	var tile GameSprite

	pos := rl.Vector2{X: float32(WorldWidth) / 2, Y: float32(WorldHeight) / 2}

	player.Init("player", "player", pos.X, pos.Y)
	slime.Init("slime", "slime_ew", pos.X, pos.Y+48)
	water.Init("water", "region1", pos.X, pos.Y+96)
	explode.Init("explode", "region5", pos.X, pos.Y+144)
	tile.Init("tile", "region1", 0.0, 0.0)
	err = tile.UpdateFrame("tile", 0.0)

	target := rl.LoadRenderTexture(WorldWidth, WorldHeight)
	defer rl.UnloadRenderTexture(target)

	//create background texture
	rl.BeginTextureMode(target)

	for x := range 32 {
		for y := range 32 {
			tile.SetX(float32(x * 48))
			tile.SetY(float32(y * 48))
			rl.DrawTextureRec(spriteSheet1, tile.Rect, tile.Pos, rl.White)
		}
	}

	rl.EndTextureMode()

	camera := rl.Camera2D{}
	var camTarget *GameSprite = &player
	//set player as target
	camera.Target = rl.NewVector2(camTarget.centerX(), camTarget.centerY())
	//set offset to camera center
	//	camera.Offset = rl.NewVector2(float32(ScreenWidth/2), float32(ScreenHeight/2))
	camera.Rotation = 0.0
	camera.Zoom = 1.0

	for !rl.WindowShouldClose() {

		dt := rl.GetFrameTime()

		accumulatedTime += dt

		if timeBetweenFrames <= accumulatedTime*gameSpeed {
			accumulatedTime = 0.0
			player.CurrentFrame++
			player.CurrentFrame = player.CurrentFrame % player.CurrentAnim.Count
			slime.CurrentFrame++
			slime.CurrentFrame = slime.CurrentFrame % slime.CurrentAnim.Count
			water.CurrentFrame++
			water.CurrentFrame = water.CurrentFrame % water.CurrentAnim.Count
			explode.CurrentFrame++
			explode.CurrentFrame = player.CurrentFrame % player.CurrentAnim.Count
		}
		err = player.UpdateFrame("walk_north", player.CurrentFrame)
		err = slime.UpdateFrame("east", slime.CurrentFrame)
		err = water.UpdateFrame("water", water.CurrentFrame)
		err = explode.UpdateFrame("explode", explode.CurrentFrame)

		camera.Target = rl.NewVector2(camTarget.centerX(), camTarget.centerY())

		rl.BeginDrawing()

		//Background
		rl.ClearBackground(rl.Black)

		t := target.Texture
		rl.DrawTextureRec(
			t,
			rl.Rectangle{X: 0, Y: 0, Width: float32(t.Width), Height: -float32(t.Height)}, // Flip vertically
			rl.Vector2{X: 0, Y: 0},
			rl.White,
		)

		//rl.BeginMode2D(camera)

		rl.DrawTextureRec(spriteSheet1, player.Rect, player.Pos, rl.White)
		rl.DrawTextureRec(spriteSheet1, water.Rect, water.Pos, rl.White)
		rl.DrawTextureRec(spriteSheet1, slime.Rect, slime.Pos, rl.White)
		if !explode.Played {
			rl.DrawTextureRec(spriteSheet1, explode.Rect, explode.Pos, rl.White)
		}
		strp := fmt.Sprintf("%v", int32(player.centerY()))
		rl.DrawText(strp, int32(player.X())-10, int32(player.Y())-20, 20, rl.White)

		//rl.EndMode2D()

		// rl.DrawText(strw, 500.0, 200.0, 40, rl.Black)
		// rl.DrawText(strs, 500.0, 300.0, 40, rl.Black)
		// if !explode.Played {
		// 	rl.DrawText(stre, 500.0, 500.0, 40, rl.Black)
		// }

		rl.DrawFPS(550, 100)

		rl.EndDrawing()

		//step upwards

		//player.StepDistance(0, 48, int(gameSpeed))

		//slime.StepDistance(1, 48, int(gameSpeed))

		if explode.CurrentAnim.Loop == false && explode.CurrentFrame == 0 {
			explode.Played = true
		}

	}

}
