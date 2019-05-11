package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	config       Config
	world        []ICell
	mutex        = &sync.Mutex{}
	wg           sync.WaitGroup
	lives        livesScores
	sim          = 0
	liveInitDome bool
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

	log("VERSION 2\n")

	for {
		start := time.Now()

		world = GeneratingNormallyDistributedWorld()
		Simulate()
		ResetWorld()

		elapsed := time.Since(start)
		fmt.Printf("\rBinomial took %s", elapsed)
	}
}
