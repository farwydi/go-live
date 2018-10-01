package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"math/rand"
)

// Описывает модель поведения живой клетки

func CreateLiveCell(x int, y int) *LiveCell {
	var cell *LiveCell = &LiveCell{cell: Cell{X: x, Y: y}}
	cell.circle = 0;
	cell.health = rand.Intn(2000)
	cell.genomeGenerator()
	return cell
}

type LiveCell struct {
	cell   Cell
	genome [64]int
	circle int // цикл по геному
	health int // жизни клетки
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
	if (e.health <= 0) {
		return false
	}
	switch e.genome[e.circle] {
		case 1: //двигаться в случайном направлении
		 e.Movie()
	}
	e.circle++
	if e.circle >= cap(e.genome) {
		e.circle = 0
	}
	e.health--
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

func (e *LiveCell) genomeGenerator() {
	for index,_ := range e.genome {
		e.genome[index] = rand.Intn(2)	// покачто геном заполняется 0 стоять, 1 идти куда попало
	}
}