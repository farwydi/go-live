package main

// Описывает модель поведения живой клетки

func CreateLiveCell() *LiveCell {

	return &LiveCell{}
}

type LiveCell struct {
}

func (l *LiveCell) Draw() {

}
