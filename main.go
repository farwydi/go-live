package main

import (
    "flag"
    "fmt"
    "math/rand"
    "os"
    "os/signal"
    "sync"
)

var (
    config       Config
    world        []ICell
    mutex        = &sync.Mutex{}
    lives        livesScores
    liveInitDome bool
    epoch        int64 = 1
    epochLog           = 0
)

// -w-width=64 -w-height=32 -w-size-cell=4 make-log=true
var (
    SeedPtr             = flag.Int64("seed", 13, "Seed")
    WidthPtr            = flag.Int("w-width", 64, "World width")
    HeightPtr           = flag.Int("w-height", 32, "World height")
    SizeCellPtr         = flag.Int("w-size-cell", 4, "World size cell")
    PrintActionPtr      = flag.Bool("make-log", true, "Print action")
    PrintActionLevelPtr = flag.Int("log-level", 0, "Print action level")
    LogFile             = flag.String("log-file", "render_engine/sim.log", "Log filename")
    GeneratorType       = flag.String("gen-type", "gwait", "Type of generator genome")
    GenomeSkip          = flag.Int64("gen-skip", 100, "Log every %n epoch")
    JumpEnable          = flag.Bool("jump", true, "genome jumper")
)

func main() {
    c := make(chan os.Signal, 1)
    signal.Notify(c)

    rand.Seed(*SeedPtr)

    config = Config{
        Width:    *WidthPtr,
        Height:   *HeightPtr,
        SizeCell: *SizeCellPtr,

        LiveMaxHealth: 100,

        EatMaxCalories: 50,

        RatingEat:     10,
        RatingMove:    5,
        RatingRecycle: 15,
    }

    fmt.Printf("%+v\n", config)

    go saverLog()

    log("VERSION 2\n")

mainLoop:
    for {
        select {
        case <-c:
            fmt.Print("\nDone.")
            workLog.Wait()
            doneLog()
            break mainLoop
        default:
            loop()

            if epoch%*GenomeSkip == 0 {
                fmt.Printf("Epoch %d save, %d\n", epoch, epochLog)
                epochLog++
            }

            epoch++
        }
    }
}

func loop() {
    world = GeneratingNormallyDistributedWorld()
    Simulate()
}
