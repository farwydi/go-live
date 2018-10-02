package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Функции для рисования

// TODO: Пока так, сперва нужно сделать систему, что бы понять, что за чем будет идти, а потом уже вводить оптимизацию
func UpdateScreen(screen *ebiten.Image) error {

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %.2f", ebiten.CurrentFPS()))

	// Цикл обработки кадра
	for index, cell := range world {

		// Если клетка умерла, то помечаем ее как удалённую
		if !cell.Action() {
			world[index] = CreateEmptyCell(calcXY(index))
		}

		cell.Draw(screen)
	}

	return nil
}
