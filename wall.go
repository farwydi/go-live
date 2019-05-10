package main

// Модель ячейка-препятствие

func CreateWellCell(x int, y int) *WellCell {

    return &WellCell{Cell{x, y}}
}

type WellCell struct {
    cell Cell
}

func (e *WellCell) Action() {
    defer wg.Done()
}
