package main

import "github.com/hajimehoshi/ebiten"

// Модель клетки, ячейки базовый объект, пустота

type ICell interface {
	Draw(screen *ebiten.Image)
	Action() bool
}

type Cell struct {
	X int
	Y int

	print *ebiten.Image
}

func (c *Cell) GetXY() (float64, float64) {
	return float64(c.X * (config.SizeCell + 1)), float64(c.Y * (config.SizeCell + 1))
}
