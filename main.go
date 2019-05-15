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

// -w-width=64 -w-height=32 -w-size-cell=4 make-log=true
var (
    SeedPtr             = flag.Int64("seed", 13, "Seed")
    WidthPtr            = flag.Int("w-width", 64, "World width")
    HeightPtr           = flag.Int("w-height", 32, "World height")
    SizeCellPtr         = flag.Int("w-size-cell", 4, "World size cell")
    PrintActionPtr      = flag.Bool("make-log", true, "Print action")
    PrintActionLevelPtr = flag.Int("log-level", 0, "Print action level")
    LogFile             = flag.String("log-file", "sim.log", "Log filename")
    GeneratorType       = flag.String("gen-type", "gwait", "Type of generator genome")
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

    eph := 1
    for {
        start := time.Now()

        loop()

        elapsed := time.Since(start)
        fmt.Printf("\rBinomial took %s, %d", elapsed, eph)
        eph++
    }
}

func loop() {
    world = GeneratingNormallyDistributedWorld()
    Simulate()
    ResetWorld()
}
