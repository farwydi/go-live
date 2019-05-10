package main

// Модель пустой ячейки

func CreateEmptyCell(x int, y int) *EmptyCell {

    return &EmptyCell{Cell{x, y}}
}

type EmptyCell struct {
    cell Cell
}

func (e *EmptyCell) Action() {
    defer wg.Done()
}
