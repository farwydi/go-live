package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"math/rand"
	"time"
)

// TODO: Пока так, сперва нужно сделать систему, что бы понять, что за чем будет идти, а потом уже вводить оптимизацию
func update(screen *ebiten.Image) error {
	for index, cell := range world {
		cell.Draw(screen)
		var cell_live = cell.Action();
		if (!cell_live) { //если клетка умерла, то помечаем ее как удаленную
			world[index] = CreateEmptyCell(calcXY(index))
		}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %.2f", ebiten.CurrentFPS()))

	// Цикл обработки кадра

	return nil
}

var config Config
var world []ICell

func main() {
	rand.Seed(time.Now().Unix())

	config = Config{
		Width:    64,
		Height:   32,
		SizeCell: 4,

		CountLiveCell:   100,
		CountPoisonCell: 64,
		CountEatCell:    32,
	}

	ww := config.Width * (config.SizeCell + 1)
	wh := config.Height * (config.SizeCell + 1)

	world = GeneratingNormallyDistributedWorld()

	ebiten.Run(update, ww, wh, 3, "Hello world!")
}
