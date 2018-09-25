package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"math/rand"
)

// TODO: Пока так, сперва нужно сделать систему, что бы понять, что за чем будет идти, а потом уже вводить оптимизацию
func update(screen *ebiten.Image) error {

	for _, cell := range world {
		cell.Draw(screen)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %.2f", ebiten.CurrentFPS()))

	// Цикл обработки кадра

	return nil
}

var config Config
var world []ICell

func main() {

	rand.Seed(13)

	config = Config{
		Width:    64,
		Height:   32,
		SizeCell: 8,

		CountLiveCell:   8,
		CountPoisonCell: 64,
		CountEatCell:    32,
	}

	ww := config.Width * (config.SizeCell + 1)
	wh := config.Height * (config.SizeCell + 1)

	world = GeneratingNormallyDistributedWorld()

	ebiten.Run(update, ww, wh, 3, "Hello world!")
}
