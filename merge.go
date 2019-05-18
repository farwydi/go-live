package main

import "math/rand"

// Механизм селекции

func inArray(b int, arr [GenomeHalfSize]int) bool {
    for _, a := range arr {
        if a == b {
            return true
        }
    }

    return false
}

func MutationStd(genome Genome) Genome {
    r := rand.Intn(1001)
    if r > 990 {
        genome[RandomGenPosition()] = RandomGen()
        log("MUTATION\n")
    }

    return genome
}

// Стондартный метод скрещивания
// Случайная точка разреза
// Честь от 1го часть от в 2го
func MergeStd(genome1, genome2 Genome) (genome Genome) {
    point := RandomGenPosition()

    for i := 0; i < GenomeSize; i++ {
        if i > point {
            genome[i] = genome2[i]
        } else {
            genome[i] = genome1[i]
        }
    }

    return
}

func MergeWTF(genome1, genome2 Genome) Genome {

    // Создаём список длиной половиной генома
    // И заполняем его случайными номерами адресов генома
    // Не повторяются
    var AGen [GenomeHalfSize]int
    for i := 0; i < GenomeHalfSize; i++ {

        NGen := RandomGenPosition()
        for inArray(NGen, AGen) {
            NGen = RandomHalfGenPosition()
        }

        AGen[i] = NGen
    }

    genome := genome1
    for _, i := range AGen {
        genome[i] = genome2[i]
    }

    genome[RandomGenPosition()] = RandomGen()

    return genome
}
