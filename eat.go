package main

// Модель ячейки с едой

func CreateEatCell() *EatCell {

	return &EatCell{}
}

type EatCell struct {
}

func (e *EatCell) Draw() {

}
