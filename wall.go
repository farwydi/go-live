package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

// Модель ячейка-препятствие

func CreateWellCell(x int, y int) *WellCell {

	return &WellCell{Cell{x, y, nil}}
}

type WellCell struct {
	cell Cell
}

func (e *WellCell) Draw(screen *ebiten.Image) {
	if e.cell.print == nil {
		e.cell.print, _ = ebiten.NewImage(config.SizeCell, config.SizeCell, ebiten.FilterNearest)
	}

	e.cell.print.Fill(color.White)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(e.cell.GetXY())

	screen.DrawImage(e.cell.print, opts)
}
