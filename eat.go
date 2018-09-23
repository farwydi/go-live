package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

// Модель ячейки с едой

func CreateEatCell(x int, y int) *EatCell {

	return &EatCell{Cell{x, y, nil}}
}

type EatCell struct {
	cell Cell
}

func (e *EatCell) Draw(screen *ebiten.Image) {
	if e.cell.print == nil {
		e.cell.print, _ = ebiten.NewImage(config.SizeCell, config.SizeCell, ebiten.FilterNearest)
	}

	e.cell.print.Fill(color.NRGBA{R: 0xff, A: 0xff})

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(e.cell.GetXY())

	screen.DrawImage(e.cell.print, opts)
}
