package main

// Модель ячейки с ядом

func CreatePoisonCell() *PoisonCell {

	return &PoisonCell{}
}

type PoisonCell struct {
}

func (e *PoisonCell) Draw() {

}
