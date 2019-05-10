package main

// Модель ячейки с ядом

func CreatePoisonCell(x int, y int) *PoisonCell {

    return &PoisonCell{Cell{x, y}}
}

type PoisonCell struct {
    cell Cell
}

func (e *PoisonCell) Action() {
    defer wg.Done()
}
