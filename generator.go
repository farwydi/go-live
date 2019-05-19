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

func AddPoisonInWorld() {
reTry:
    i := rand.Intn(config.Height * config.Width)
    switch world[i].(type) {
    case *EmptyCell:
        world[i] = CreatePoisonCell(calcXY(i))
    default:
        goto reTry
    }
}

func AddEatInWorld() {
reTry:
    i := rand.Intn(config.Height * config.Width)
    switch world[i].(type) {
    case *EmptyCell:
        world[i] = CreateEatCell(calcXY(i))
    default:
        goto reTry
    }
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

    // Custom
    lives[0].name = "alpha"
    lives[0].genome = Genome{
        GSeeRight,
        GEnd + 17,
        GEnd + 8,  // EatCell
        GEnd + 11, // PoisonCell
        GEnd + 15, // EmptyCell
        GEnd + 17, // WellCell
        GEnd + 17, // LiveCell
        GEatRight,
        GMoveRight,
        GEnd + 1,
        GRecycleRight,
        GEatRight,
        GMoveRight,
        GEnd + 1,
        GMoveRight,
        GEnd + 1,
        GEnd + 18,
        GSeeUp,
        GEnd + 34,
        GEnd + 25, // EatCell
        GEnd + 28, // PoisonCell
        GEnd + 32, // EmptyCell
        GEnd + 34, // WellCell
        GEnd + 34, // LiveCell
        GEatUp,
        GMoveUp,
        GEnd + 18,
        GRecycleUp,
        GEatUp,
        GMoveUp,
        GEnd + 18,
        GMoveUp,
        GEnd + 18,
        GEnd + 35,
        GSeeLeft,
        GEnd + 51,
        GEnd + 42, // EatCell
        GEnd + 45, // PoisonCell
        GEnd + 49, // EmptyCell
        GEnd + 51, // WellCell
        GEnd + 51, // LiveCell
        GEatLeft,
        GMoveLeft,
        GEnd + 35,
        GRecycleLeft,
        GEatLeft,
        GMoveLeft,
        GEnd + 35,
        GMoveLeft,
        GEnd + 35,
        GEnd + 52,
        GSeeDown,
        GEnd + 68,
        GEnd + 59, // EatCell
        GEnd + 62, // PoisonCell
        GEnd + 66, // EmptyCell
        GEnd + 68, // WellCell
        GEnd + 68, // LiveCell
        GEatDown,
        GMoveDown,
        GEnd + 52,
        GRecycleDown,
        GEatDown,
        GMoveDown,
        GEnd + 52,
        GMoveDown,
        GEnd + 52,
        GEnd + 1,
    }

    lives[1].name = "beta"
    lives[1].genome = lives[0].genome

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
