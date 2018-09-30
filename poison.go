package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

// Модель ячейки с ядом

func CreatePoisonCell(x int, y int) *PoisonCell {

	return &PoisonCell{Cell{x, y, nil}}
}

type PoisonCell struct {
	cell Cell
}

func (e *PoisonCell) Draw(screen *ebiten.Image) {

	if e.cell.print == nil {
		e.cell.print, _ = ebiten.NewImage(config.SizeCell, config.SizeCell, ebiten.FilterNearest)
	}

	e.cell.print.Fill(color.NRGBA{G: 0xff, A: 0xff})

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(e.cell.GetXY())

	screen.DrawImage(e.cell.print, opts)
}

func (e *PoisonCell) Action() bool {
	return true;
}