package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

// Описывает модель поведения живой клетки

func CreateLiveCell(x int, y int) *LiveCell {

	return &LiveCell{Cell{x, y, nil}}
}

type LiveCell struct {
	cell Cell
}

func (e *LiveCell) Draw(screen *ebiten.Image) {

	if e.cell.print == nil {
		e.cell.print, _ = ebiten.NewImage(config.SizeCell, config.SizeCell, ebiten.FilterNearest)
	}

	e.cell.print.Fill(color.NRGBA{G: 0xff, R: 0xff, A: 0xff})

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(e.cell.GetXY())

	screen.DrawImage(e.cell.print, opts)
}
