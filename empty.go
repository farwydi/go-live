package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

// Модель пустой ячейки

func CreateEmptyCell(x int, y int) *EmptyCell {

	return &EmptyCell{Cell{x, y, nil}}
}

type EmptyCell struct {
	cell Cell
}

func (e *EmptyCell) Draw(screen *ebiten.Image) {
	if e.cell.print == nil {
		e.cell.print, _ = ebiten.NewImage(config.SizeCell, config.SizeCell, ebiten.FilterNearest)
	}

	e.cell.print.Fill(color.Black)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(e.cell.GetXY())

	screen.DrawImage(e.cell.print, opts)
}
