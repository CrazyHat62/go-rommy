# GO ROMMY

This is an exercise in writing a game, building a library for sprite atlases, 
and learning Go. The spriteAtlas could be a reader but chose to do it this way
because everything is still fuzzy.  

## SpriteAtlas

A simple library to read a simple atlas for a simple sprite-sheet.
This Atlas is simple on purpose and does not apply many restrictions
The Atlas could be considered Alpha as the design can be further 
simplified at any time.

## Using the library

in the go.mod file :

    github.com/CrazyHat62/SpriteAtlas v0.1.1

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

To use sprite atlas during development (and make changes) we simply made spriteAtlas a submodule, but this doesn't need to be the case as it is a seperate
go repository and can be imported as usual 

    replace spriteatlas => ../spriteatlas // Relative path for submodule found in mod

