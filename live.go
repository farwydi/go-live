package main

import (
	"fmt"
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
		cell: Cell{x, y},
		// Параметр клетки, по умолчанию равен максимальному значению
		health: config.LiveMaxHealth,
		name:   RandStringRunes(5),
	}
}

type livesScores [CountLiveCell]*LiveCell

func (s livesScores) Len() int {
	return len(s)
}

func (s livesScores) Swap(i, j int) {
	*s[i], *s[j] = *s[j], *s[i]
}

func (s livesScores) Less(i, j int) bool {
	return s[i].score > s[j].score
}

type LiveCell struct {
	cell   Cell
	genome Genome
	health int // жизни клетки
	score  int // рейтинг клетки
	name   string
}

func (e *LiveCell) IsLive() bool {
	if e.health <= 0 {
		// Означает что клетка умерла :(
		return false
	}

	return true
}

func (e *LiveCell) See(vector [2]int) {

}

func (e *LiveCell) Movie(vector [2]int) error {

	// movieX, movieY - координаты движения
	// i - адрес в массиве
	movieX := e.cell.X + vector[0]
	movieY := e.cell.Y - vector[1]
	i, err := resolveXY(movieX, movieY)

	// Движение за границы
	if err != nil {
		return err
	}

	switch t := world[i].(type) {
	case *EmptyCell:
		// Наткнулись на пустую клетку
		if PrintAction {
			fmt.Printf("[%s.%d] move (%d,%d)\n", e.name, e.health, movieX, movieY)
		}
		mutex.Lock()
		e.cell.X = movieX
		e.cell.Y = movieY
		mutex.Unlock()
		e.score += config.RatingMove

	case *PoisonCell:
		// Наступили на яд
		if PrintAction {
			fmt.Printf("[%s.%d] move and die (poison)\n", e.name, e.health)
		}
		e.health = 0
	case *EatCell:
		// Наступили на еду
		if PrintAction {
			fmt.Printf("[%s.%d] move and eat\n", e.name, e.health)
		}
		// Забираем калории еды и добавляем их к текущему ХП
		e.health += t.calories
		// Начисляем рейтинг за активность
		e.score += config.RatingEat
		// Обнуляет клетку с едой
		world[i] = CreateEmptyCell(movieX, movieY)
		// Переходим на эту клетку

		// Двигаем клетку на место с едой
		mutex.Lock()
		e.cell.X = movieX
		e.cell.Y = movieY
		mutex.Unlock()
	}

	// ОК
	return nil
}

func RandGenomeGenerator() Genome {
	var genome Genome
	for index := range genome {
		genome[index] = RandomGen() // пока что геном заполняется рандомно
	}
	return genome
}

// Команды генома:
const (
	GWait = iota

	// Движение
	GMoveUp
	GMoveUpLeft
	GMoveUpRight
	GMoveLeft
	GMoveRight
	GMoveDown
	GMoveDownLeft
	GMoveDownRight

	// Посмотреть
	GSeeUp
	GSeeUpLeft
	GSeeUpRight
	GSeeLeft
	GSeeRight
	GSeeDown
	GSeeDownLeft
	GSeeDownRight

	// Конец команд
	GEnd

	GJumpStart
	GJumpEnd = GJumpStart + (GEnd - 1)
)

func (e *LiveCell) Action() {
	defer wg.Done()

	// Цикл крутится пока не закончатся ХП
	for it := 0; e.IsLive(); e.health-- {

		// Выполняем команду в геноме
		switch e.genome[it] {

		case GWait:
			// Ничего не делать
			if PrintAction && PrintActionLevel > 2 {
				fmt.Printf("[%s.%d] wait\n", e.name, e.health)
			}

		case GMoveUp:
			e.Movie(Up)
		case GMoveUpRight:
			e.Movie(UpRight)
		case GMoveRight:
			e.Movie(Right)
		case GMoveDownRight:
			e.Movie(DownRight)
		case GMoveDown:
			e.Movie(Down)
		case GMoveDownLeft:
			e.Movie(DownLeft)
		case GMoveLeft:
			e.Movie(Left)
		case GMoveUpLeft:
			e.Movie(UpLeft)

		default:
			jumpTo := int(e.genome[it])
			switch {
			// Это номер в геноме куда переместить указатель
			case jumpTo >= GJumpStart && jumpTo <= GJumpEnd:
				// Безусловный переход
				if PrintAction && PrintActionLevel > 2 {
					fmt.Printf("[%s.%d] seek %d\n", e.name, e.health, jumpTo)
				}
				it = jumpTo - (GEnd + 1)
			default:
				// Неизвестная команда
				it++
			}
		}

		// Зацикливание
		if it > GenomeSize-1 {
			it = 0
		}
	}
}
