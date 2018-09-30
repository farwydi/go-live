package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

// Описывает модель поведения живой клетки

func CreateLiveCell(x int, y int) *LiveCell {

	return &LiveCell{cell: Cell{X: x, Y: y}}
}

type LiveCell struct {
	cell Cell
	genome [64]int
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
