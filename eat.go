package main

import (
	"fmt"
	"math/rand"
)

// Модель ячейки с едой

func CreateEatCell(x int, y int) *EatCell {

	log(fmt.Sprintf("I E %d,%d\n", x, y))

	return &EatCell{
		cell:     Cell{x, y},
		calories: rand.Intn(config.EatMaxCalories),
	}
}

type EatCell struct {
	cell     Cell
	calories int
}

func (e *EatCell) Action() {
	defer wg.Done()
}
