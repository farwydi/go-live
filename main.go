package main

import (
    "flag"
    "fmt"
    "math/rand"
    "sync"
    "time"
)

var (
    config       Config
    world        []ICell
    mutex        = &sync.Mutex{}
    wg           sync.WaitGroup
    lives        livesScores
    sim          = 0
    liveInitDome bool
)

// -wWidth=64 -wHeight=32 -wSizeCell=4
var (
    SeedPtr             = flag.Int64("seed", 13, "Seed")
    WidthPtr            = flag.Int("wWidth", 64, "World width")
    HeightPtr           = flag.Int("wHeight", 32, "World height")
    SizeCellPtr         = flag.Int("wSizeCell", 4, "World size cell")
    PrintActionPtr      = flag.Bool("action", false, "Print action")
    PrintActionLevelPtr = flag.Int("wActionLevel", 3, "Print action level")
)

func main() {
    rand.Seed(*SeedPtr)

    config = Config{
        Width:    *WidthPtr,
        Height:   *HeightPtr,
        SizeCell: *SizeCellPtr,

        LiveMaxHealth: 100,

        EatMaxCalories: 50,

        RatingEat:  10,
        RatingMove: 5,
    }

    fmt.Printf("%+v\n", config)

    for {
        start := time.Now()
        loop()
        fmt.Printf("\rBinomial took %s", time.Since(start))
    }
}

func loop() {
    world = GeneratingNormallyDistributedWorld()
    Simulate()
    ResetWorld()

    //ebiten.Run(UpdateScreen, ww, wh, 3, "Hello world!")
}

//func UpdateScreen(screen *ebiten.Image) error {
//
//    ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %.2f", ebiten.CurrentFPS()))
//
//    // Цикл обработки кадра
//    for index, cell := range world {
//
//        // Если клетка умерла, то помечаем ее как удалённую
//        if !cell.Action() {
//            world[index] = CreateEmptyCell(calcXY(index))
//        }
//
//        cell.Draw(screen)
//    }
//
//    return nil
//}
