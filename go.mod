module Rommy2

go 1.25.2

replace github.com/CrazyHat62/spriteatlas => ./libs/spriteatlas // Relative path example

require (
	github.com/CrazyHat62/spriteatlas v0.0.0-20251029204420-80c95cb9b3b0
	github.com/gen2brain/raylib-go/raylib v0.55.1
	golang.org/x/image v0.32.0
)

require (
	github.com/ebitengine/purego v0.7.1 // indirect
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
	golang.org/x/sys v0.20.0 // indirect
)
