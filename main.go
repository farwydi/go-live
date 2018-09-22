package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// TODO: Пока так, сперва нужно сделать систему, что бы понять, что за чем будет идти, а потом уже вводить оптимизацию
func update(screen *ebiten.Image) error {
	ebitenutil.DebugPrint(screen, "Hello world!")

	// Цикл обработки кадра

	return nil
}

var config Config
var world []Cell

func main() {

	config = Config{
		120,
		120,
		8,
	}

	ww := config.Height * (config.SizeCell + 1)
	wh := config.Width * (config.SizeCell + 1)

	world = GeneratingNormallyDistributedWorld(config.Height, config.Width)

	ebiten.Run(update, ww, wh, 1, "Hello world!")
}
