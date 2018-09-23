package main

import "github.com/hajimehoshi/ebiten"

// Описывает модель поведения живой клетки

func CreateLiveCell(x int, y int) *LiveCell {

	return &LiveCell{Cell{x, y, nil}}
}

type LiveCell struct {
	cell Cell
}

func (l *LiveCell) Draw(screen *ebiten.Image) {

}
