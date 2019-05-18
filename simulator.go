package main

import (
    "sort"
    "strconv"
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

        //log("SHOT\n")

        for _, cell := range world {
            switch t := cell.(type) {
            case *LiveCell:
                if !t.Action() {
                    running = true
                }
            }
        }
    }

    // Селекция
    sort.Sort(lives)

    //leader := lives[:CountLiveCell/2]

    for i := 0; i < CountLiveCell/2; i++ {
        lives[i+CountLiveCell/2].genome = Merge(lives[i].genome, lives[i+1].genome)
    }

    s := "\r"

    for i := 0; i < 10; i++ {
        s += strconv.Itoa(lives[i].score) + " "
    }

    //fmt.Print(s)
    liveInitDome = true
}
