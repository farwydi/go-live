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

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %.2f", ebiten.CurrentFPS()))

	// Цикл обработки кадра
	for index, cell := range world {

		// Если клетка умерла, то помечаем ее как удалённую
		if !cell.Action() {
			world[index] = CreateEmptyCell(calcXY(index))
		}

		cell.Draw(screen)
	}

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

		LiveMaxHealth: 2000,
		LiveMaxThing:  10,

		EatMaxCalories: 50,

		Rating: map[string]int{
			"eat": 10,
			"movie": 5,
		},
	}

	ww := config.Width * (config.SizeCell + 1)
	wh := config.Height * (config.SizeCell + 1)

	world = GeneratingNormallyDistributedWorld()

	ebiten.Run(update, ww, wh, 3, "Hello world!")
}
