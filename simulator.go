package main

// Функции запуска симуляции

func Simulate() {

	if world != nil && len(world) == 0 {
		panic("world not init")
	}

	// Цикл обработки кадра
	for _, cell := range world {

		// Отпускаем клетку думать
		wg.Add(1)
		go cell.Action()
	}

	wg.Wait()

	// Селекция
	//for _, live := range lives {
	//live
	//}
}
