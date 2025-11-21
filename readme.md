# GO ROMMY

This is an exercise in writing a game, building a library for sprite atlases, 
and learning Go. The SpriteAtlas could be a reader but chose to do it this way
because everything is still fuzzy.  

## SpriteAtlas

A simple library to read a simple atlas for a simple sprite-sheet.
This Atlas is simple on purpose and does not apply many restrictions
The Atlas could be considered Alpha as the design can be further 
simplified at any time.

## Using the library

in the go.mod file :

    github.com/CrazyHat62/SpriteAtlas v0.1.2

in your main.go

```
import (
	"fmt"
	"os"

	sa "github.com/CrazyHat62/SpriteAtlas"
	rl "github.com/gen2brain/raylib-go/raylib"
)
```

## Go Mod

To use sprite atlas during development (and make changes) we simply made SpriteAtlas a submodule, but this doesn't need to be the case as it is a separate
go repository and can be imported as usual 

// Relative path example
replace github.com/CrazyHat62/SpriteAtlas => ./libs/spriteatlas 

NOTE:
According to the web. 
Tile images: The tileset, created by Anders Kaseorg, has been explicitly placed in the public domain. It is not covered by the MIT licence.

