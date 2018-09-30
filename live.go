package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"math/rand"
)

// Описывает модель поведения живой клетки

func CreateLiveCell(x int, y int) *LiveCell {

	return &LiveCell{cell: Cell{X: x, Y: y}}
}

type LiveCell struct {
	cell   Cell
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

func (e *LiveCell) Action() bool {

	e.Movie()
	return true
}

func (e *LiveCell) Movie() {
	for {
		directionVector := rand.Intn(8) + 1
		var movieX int
		var movieY int
		switch directionVector {
		case 1:
			movieX = e.cell.X - 1
			movieY = e.cell.Y + 1
		case 2:
			movieX = e.cell.X
			movieY = e.cell.Y + 1
		case 3:
			movieX = e.cell.X + 1
			movieY = e.cell.Y + 1
		case 4:
			movieX = e.cell.X + 1
			movieY = e.cell.Y
		case 5:
			movieX = e.cell.X + 1
			movieY = e.cell.Y - 1
		case 6:
			movieX = e.cell.X
			movieY = e.cell.Y - 1
		case 7:
			movieX = e.cell.X - 1
			movieY = e.cell.Y + 1
		case 8:
			movieX = e.cell.X - 1
			movieY = e.cell.Y
		}

		size := len(world)
		index := (movieX * config.Height) + movieY

		if index > size || index < 0 {
			continue
		}

		switch world[index].(type) {
		case *WellCell:
			continue
		}

		e.cell.X = movieX
		e.cell.Y = movieY
		break
	}
}
