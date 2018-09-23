package main

// Механизм генерации мира
// Модель:
// y
// ^
// |
// |
// |
// |(x, y)
// .----------> x

// Простая функция создания нормально распределённого мира
func GeneratingNormallyDistributedWorld(countByX int, countByY int) []ICell {

	size := countByX * countByY

	world := make([]ICell, size)

	for i := 0; i < size; {

		x := i / countByX
		y := i - (x * countByX)

		// TODO: Добавить хаоса
		world[i] = CreateLiveCell(x, y)
		world[i+1] = CreateEatCell(x, y+1)
		world[i+2] = CreatePoisonCell(x, y+2)

		// TODO: Поменять
		world[i+3] = CreateWellCell(x, y+3)

		i = i + 4
	}

	return world
}
