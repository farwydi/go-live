package main

import (
    "fmt"
    "sort"
)

// Функции запуска симуляции

func Simulate() {

    if world != nil && len(world) == 0 {
        panic("world not init")
    }

    for _, cell := range world {
        switch t := cell.(type) {
        case *LiveCell:
            t.PreviewGenome()
        }
    }

    // Цикл обработки кадра
    running := true
    for running {
        running = false

        for _, cell := range world {
            switch t := cell.(type) {
            case *LiveCell:
                if !t.Action() {
                    running = true
                }
            }
        }

        log("SHOT_END\n")
    }

    // Селекция
    sort.Sort(lives)

    //leader := lives[:CountLiveCell/2]

    for i := 0; i < CountLiveCell/2; i++ {
        lives[i+CountLiveCell/2].genome = MutationStd(MergeStd(lives[i].genome, lives[i+1].genome))
    }

    fmt.Printf("\r%d\t", lives[0].score)
    liveInitDome = true
}
