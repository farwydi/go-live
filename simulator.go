package main

import "fmt"

// Функции запуска симуляции

func Simulate() {

	if world != nil && len(world) == 0 {
		panic("world not init")
	}

	//for {
	// Цикл обработки кадра
	for index, cell := range world {

		live, ok := cell.(*LiveCell)
		if ok {
			c := make(chan bool)

			// Отпускаем клетку думать
			go live.Detach(c)

			// Габелла
			if !<-c {
				fmt.Printf("[%s] done", live.name)
				world[index] = CreateEmptyCell(calcXY(index))
			}
		}
	}
	//}
}
