package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"math/rand"
	"os"
	"sync"
)

var (
	config Config
	world  []ICell
	mutex  = &sync.Mutex{}
)

func main() {
	rand.Seed(13)

	config = Config{
		Width:    64,
		Height:   32,
		SizeCell: 4,

		CountLiveCell:   100,
		CountPoisonCell: 64,
		CountEatCell:    32,

		LiveMaxHealth: 100,

		EatMaxCalories: 50,

		RatingEat:  10,
		RatingMove: 5,
	}

	fmt.Printf("%v+\n", config)

	world = GeneratingNormallyDistributedWorld()

	if os.Getenv("simulate") == "on" {

		Simulate()
	} else {

		ww := config.Width * (config.SizeCell + 1)
		wh := config.Height * (config.SizeCell + 1)
		ebiten.Run(UpdateScreen, ww, wh, 3, "Hello world!")
	}
}
