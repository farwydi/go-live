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

		// 10% Well
		if InBetween(r, 0, 10) {
			world[i] = CreateWellCell(x, y)
		}

		// 30% Poison
		if InBetween(r, 11, 40) {
			world[i] = CreatePoisonCell(x, y)
		}

		// 40% Eat
		if InBetween(r, 41, 80) {
			world[i] = CreateEatCell(x, y)
		}

		// 20% Live
		if InBetween(r, 81, 100) {
			world[i] = CreateLiveCell(x, y)
		}
	}

	return world
}
