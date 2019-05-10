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

// -wwidth=64 -wheight=32 -wsizecell=4
var (
    WidthPtr    = flag.Int("wwidth", 64, "World width")
    HeightPtr   = flag.Int("wheight", 32, "World height")
    SizeCellPtr = flag.Int("wsizecell", 4, "World size cell")
)

func main() {
    rand.Seed(13)

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
}
