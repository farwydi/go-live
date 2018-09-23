package main

import "math/rand"

// Механизм генерации мира
// Модель:
// y
// ^
// |
// |
// |
// |(x, y)
// .----------> x

func InBetween(i, min, max int32) bool {
	return (i >= min) && (i <= max)
}

// Простая функция создания нормально распределённого мира
func GeneratingNormallyDistributedWorld(countByX int, countByY int) []ICell {

	size := countByX * countByY

	world := make([]ICell, size)

	for i := 0; i < size; i++ {

		x := i / countByX
		y := i - (x * countByX)

		r := rand.Int31n(100)

		// 2% Well
		if InBetween(r, 0, 2) {
			world[i] = CreateWellCell(x, y)
		}

		// 4% Poison
		if InBetween(r, 3, 7) {
			world[i] = CreatePoisonCell(x, y)
		}

		// 5% Eat
		if InBetween(r, 8, 13) {
			world[i] = CreateEatCell(x, y)
		}

		// 1% Live
		if InBetween(r, 14, 15) {
			world[i] = CreateLiveCell(x, y)
		}

		// 88% Empty
		if InBetween(r, 16, 100) {
			world[i] = CreateEmptyCell(x, y)
		}
	}

	return world
}
