package main

// Настройки

type Config struct {
	// Размерность мира, кол-во клеток в направлениях x и y
	Width  int
	Height int

	// Размер 1ой клетки в пикселях
	SizeCell int

	CountLiveCell   int
	CountPoisonCell int
	CountEatCell    int

	// Параметры живой клетки
	LiveMaxHealth int
	LiveMaxThing  int
}
