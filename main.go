package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func update(screen *ebiten.Image) error {
	ebitenutil.DebugPrint(screen, "Hello world!")
	return nil
}

var config Config

func main() {

	config = Config{
		120,
		120,
		8,
	}

	ww := config.Height * (config.SizeCell + 1)
	wh := config.Width * (config.SizeCell + 1)

	ebiten.Run(update, ww, wh, 1, "Hello world!")
}
