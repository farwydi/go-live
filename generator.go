package main

// Механизм генерации мира

// Простая функция создания нормально распределённого мира
func GeneratingNormallyDistributedWorld(countByX int, countByY int) []Cell {

	size := countByX * countByY

	world := make([]Cell, size)

	for i := 0; i < size; {
		// TODO: Добавить хаоса
		world[i] = CreateLiveCell()
		world[i+1] = CreateEatCell()
		world[i+2] = CreatePoisonCell()

		// TODO: Поменять
		world[i+3] = CreateWellCell()

		i = i + 4
	}

	return world
}
