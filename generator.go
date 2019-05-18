package main

import (
    "errors"
    "fmt"
    "math/rand"
)

// Механизм генерации мира
// Модель:
// y
// ^
// |
// |
// |
// |(x, y)
// .----------> x

// Получить X и Y по индексу
func calcXY(i int) (int, int) {

    if i > config.Height*config.Width {
        panic("Out of range")
    }

    x := i / config.Height
    y := i - (x * config.Height)

    return x, y
}

// Получить индекс в массиве по X и Y
func resolveXY(x int, y int) (int, error) {

    if x > config.Width {
        return 0, errors.New("Y > Width")
    }

    if y > config.Height {
        return 0, errors.New("X > Height")
    }

    return (config.Height * x) + y, nil
}

// Простая функция создания нормально распределённого мира
func GeneratingNormallyDistributedWorld() []ICell {

    if config.Height == 0 {
        panic("Height zero")
    }

    if config.Width == 0 {
        panic("Width zero")
    }

    size := config.Height * config.Width
    world := make([]ICell, size)

    log(fmt.Sprintf("EPOCH %d\n", epochLog))

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
    liveIt := 0
    for c := CountLiveCell; c > 0; {

        i := rand.Intn(size)

        if world[i] == nil {
            x, y := calcXY(i)
            live := CreateLiveCell(x, y, liveIt)

            if liveInitDome {
                live.genome = lives[liveIt].genome
            } else {
                switch *GeneratorType {
                case "gwait":
                    live.genome = GWaitGenomeGenerator()
                case "rand":
                    live.genome = RandGenomeGenerator()
                }
            }

            lives[liveIt] = live
            world[i] = live

            liveIt++
            c--
        }
    }

    // Eat
    for c := CountEatCell; c > 0; {

        i := rand.Intn(size)

        if world[i] == nil {
            world[i] = CreateEatCell(calcXY(i))
            c--
        }
    }

    // Poison
    for c := CountPoisonCell; c > 0; {

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
