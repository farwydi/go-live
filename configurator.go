package main

import "runtime"

// Настройки

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

const (
	CountLiveCell   = 100
	CountPoisonCell = 64
	CountEatCell    = 32

	PrintAction      = false
	PrintActionLevel = 3

	PrintLog = false
)

type Config struct {
	// Размерность мира, кол-во клеток в направлениях x и y
	Width  int
	Height int

	// Размер 1ой клетки в пикселях
	SizeCell int

	// Параметры живой клетки
	LiveMaxHealth int
	LiveMaxThing  int

	// Параметры еды
	EatMaxCalories int

	// Очки, начисляемые за действия
	RatingEat  int
	RatingMove int
}
