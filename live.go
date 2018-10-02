package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"math/rand"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// Описывает модель поведения живой клетки

func CreateLiveCell(x int, y int) *LiveCell {

	return &LiveCell{
		cell: Cell{X: x, Y: y},
		// Параметр клетки, по умолчанию равен максимальному значению
		health: config.LiveMaxHealth,
		name:   RandStringRunes(5),
	}
}

type LiveCell struct {
	cell   Cell
	genome [64]int
	health int // жизни клетки
	score  int // рейтинг клетки
	name   string
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

	i := 0
	for ; e.IsLive(); e.health-- {

		switch e.genome[i] {
		case 0, 66:
			// Ничего не делать
			// Выход
			fmt.Printf("[%s.%d] wait\n", e.name, e.health)
			// Движение
		case 66 + 1:
			// Верх
			e.Movie([2]int{0, 1})
		case 66 + 2:
			// Верх-право
			e.Movie([2]int{1, 1})
		case 66 + 3:
			// Право
			e.Movie([2]int{1, 0})
		case 66 + 4:
			// Низ-право
			e.Movie([2]int{1, -1})
		case 66 + 5:
			// Низ
			e.Movie([2]int{0, -1})
		case 66 + 6:
			// Низ-лево
			e.Movie([2]int{-1, -1})
		case 66 + 7:
			// Лево
			e.Movie([2]int{-1, 0})
		case 66 + 8:
			// Верх-лево
			e.Movie([2]int{-1, 1})
		}

		// Безусловный переход
		switch e.genome[i] {
		case 1, 2, 3, 4, 5, 6, 7, 8, 9,
			10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
			20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
			30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
			40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
			50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
			60, 61, 62, 63, 64: // 0-63
			i = e.genome[i] - 1
			fmt.Printf("[%s.%d] seek %d\n", e.name, e.health, i)
		default:
			i++
		}

		if i > 63 {
			i = 0
		}
	}

	return false
}

func (e *LiveCell) IsLive() bool {
	if e.health <= 0 {
		// Означает что клетка умерла :(
		return false
	}

	return true
}

func (e *LiveCell) Movie(vector [2]int) error {

	movieX := e.cell.X + vector[0]
	movieY := e.cell.Y - vector[1]
	i, err := resolveXY(movieX, movieY)

	if err != nil {
		return err
	}

	switch world[i].(type) {
	case *EmptyCell:
		fmt.Printf("[%s.%d] move (%d,%d)\n", e.name, e.health, movieX, movieY)
		mutex.Lock()
		e.cell.X = movieX
		e.cell.Y = movieY
		mutex.Unlock()
		e.score += config.RatingMove
	case *PoisonCell:
		fmt.Printf("[%s.%d] move and die (poison)\n", e.name, e.health)
		e.health = 0
	case *EatCell:
		fmt.Printf("[%s.%d] move and eat\n", e.name, e.health)
		e.health += world[i].(*EatCell).calories
		e.score += config.RatingEat
		world[i] = CreateEmptyCell(movieX, movieY)
	}

	return nil
}

func (e *LiveCell) RandGenomeGenerator() {
	for index := range e.genome {
		e.genome[index] = rand.Intn(74) // пока что геном заполняется рандомно
	}
}

func (e *LiveCell) Detach(c chan bool) {

	// Передаём статус клетки каналу
	c <- e.Action()
}
