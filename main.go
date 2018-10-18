package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var (
	config Config
	world  []ICell
	mutex  = &sync.Mutex{}
	wg     sync.WaitGroup
	lives  livesScores
)

func main() {
	rand.Seed(13)

	config = Config{
		Width:    64,
		Height:   32,
		SizeCell: 4,

		LiveMaxHealth: 100,

		EatMaxCalories: 50,

		RatingEat:  10,
		RatingMove: 5,
	}

	fmt.Printf("%+v\n", config)

	for {
		world = GeneratingNormallyDistributedWorld()
		Simulate()
		ResetWorld()
	}
}
