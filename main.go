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

    log("VERSION 2\n")

    for {
        start := time.Now()

        loop()

        elapsed := time.Since(start)
        fmt.Printf("\rBinomial took %s", elapsed)
    }
}

func loop() {
    world = GeneratingNormallyDistributedWorld()
    Simulate()
    ResetWorld()
}
