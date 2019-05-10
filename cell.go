package main

// Модель клетки, ячейки базовый объект, пустота

type ICell interface {
    Action()
}

type Cell struct {
    X int
    Y int
}

func (c *Cell) GetXY() (float64, float64) {
    return float64(c.X * (config.SizeCell + 1)), float64(c.Y * (config.SizeCell + 1))
}
