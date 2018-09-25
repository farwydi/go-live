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

//func inBetween(i, min, max int32) bool {
//	return (i >= min) && (i <= max)
//}

func calcXY(i int) (int, int) {
	x := i / config.Height
	y := i - (x * config.Height)
	return x, y
}

// Простая функция создания нормально распределённого мира
func GeneratingNormallyDistributedWorld() []ICell {

	size := config.Height * config.Width
	world := make([]ICell, size)

	// Gen well
	// Top
	for i := 0; i < size; i += config.Height {

		world[i] = CreateWellCell(calcXY(i))
	}

	// Bottom
	for i := config.Height - 1; i < size; i += config.Height {

		world[i] = CreateWellCell(calcXY(i))
	}

	// Left
	for i := 1; i < config.Height-1; i++ {

		world[i] = CreateWellCell(calcXY(i))
	}

	// Right
	for i := config.Width*config.Height - config.Height + 1; i < size-1; i++ {

		world[i] = CreateWellCell(calcXY(i))
	}

	// Live
	for c := config.CountLiveCell; c > 0; {

		i := rand.Intn(size)

		if world[i] == nil {
			world[i] = CreateLiveCell(calcXY(i))
			c--
		}
	}

	// Eat
	for c := config.CountEatCell; c > 0; {

		i := rand.Intn(size)

		if world[i] == nil {
			world[i] = CreateEatCell(calcXY(i))
			c--
		}
	}

	// Poison
	for c := config.CountPoisonCell; c > 0; {

		i := rand.Intn(size)

		if world[i] == nil {
			world[i] = CreatePoisonCell(calcXY(i))
			c--
		}
	}

	// Заполнение пустотой
	for i := 0; i < size; i++ {

		if world[i] == nil {
			world[i] = CreateEmptyCell(calcXY(i))
		}
	}

	return world
}
