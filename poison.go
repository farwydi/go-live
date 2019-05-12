package main

import "fmt"

// Модель ячейки с ядом

func CreatePoisonCell(x int, y int) *PoisonCell {

    log(fmt.Sprintf("I P %d,%d\n", x, y))

    return &PoisonCell{Cell{x, y}}
}

type PoisonCell struct {
    cell Cell
}

func (e *PoisonCell) Action() {
    defer wg.Done()
}
