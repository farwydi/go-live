package main

import (
	"github.com/hajimehoshi/ebiten"
)

// Модель пустой ячейки

func CreateEmptyCell(x int, y int) *EmptyCell {

	return &EmptyCell{Cell{x, y, nil}}
}

type EmptyCell struct {
	cell Cell
}

func (e *EmptyCell) Draw(screen *ebiten.Image) {

	// Рисовать не нужно
}

func (e *EmptyCell) Action() bool{
	return true;
}
