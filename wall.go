package main

import "fmt"

// Модель ячейка-препятствие

func CreateWellCell(x int, y int) *WellCell {

    log(fmt.Sprintf("I W %d,%d\n", x, y))

    return &WellCell{Cell{x, y}}
}

type WellCell struct {
    cell Cell
}

func (e *WellCell) Action() {
    defer wg.Done()
}
