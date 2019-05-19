package main

import (
    "fmt"
    "math/rand"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

// Описывает модель поведения живой клетки

func CreateLiveCellWithGenome(x int, y int, id int, genome Genome) *LiveCell {

    log(fmt.Sprintf("I L %d,%d\n", x, y))

    return &LiveCell{
        cell: Cell{x, y},
        // Параметр клетки, по умолчанию равен максимальному значению
        health: config.LiveMaxHealth,
        name:   RandStringRunes(5),
        id:     id,
        genome: genome,
    }
}

func CreateLiveCell(x int, y int, id int) *LiveCell {

    log(fmt.Sprintf("I L %d,%d\n", x, y))

    return &LiveCell{
        cell: Cell{x, y},
        // Параметр клетки, по умолчанию равен максимальному значению
        health: config.LiveMaxHealth,
        name:   RandStringRunes(5),
        id:     id,
    }
}

type livesScores [CountLiveCell]*LiveCell

func (s livesScores) Len() int {
    return len(s)
}

func (s livesScores) Swap(i, j int) {
    *s[i], *s[j] = *s[j], *s[i]
}

func (s livesScores) Less(i, j int) bool {
    return s[i].score > s[j].score
}

type LiveCell struct {
    cell   Cell
    genome Genome
    health int // жизни клетки
    score  int // рейтинг клетки
    name   string
    id     int
    it     int
}

func (e *LiveCell) IsLive() bool {
    if e.health <= 0 {
        // Означает что клетка умерла :(
        return false
    }

    return true
}

func (e *LiveCell) See(vector [2]int) int {
    X := e.cell.X + vector[0]
    Y := e.cell.Y - vector[1]

    i, err := resolveXY(X, Y)
    if err != nil {
        // Идём дальше
        return 1
    }

    switch world[i].(type) {
    case *EatCell:
        //if e.name == "alpha" {
        //    fmt.Printf("See EatCell\n", )
        //}
        return 2
    case *PoisonCell:
        //if e.name == "alpha" {
        //    fmt.Printf("See PoisonCell\n", )
        //}
        return 3
    case *EmptyCell:
        //if e.name == "alpha" {
        //    fmt.Printf("See EmptyCell\n", )
        //}
        return 4
    case *WellCell:
        //if e.name == "alpha" {
        //    fmt.Printf("See WellCell\n", )
        //}
        return 5
    case *LiveCell:
        //if e.name == "alpha" {
        //    fmt.Printf("See LiveCell\n")
        //}
        return 6
    }

    return 1
}

func (e *LiveCell) Eat(vector [2]int) int {
    X := e.cell.X + vector[0]
    Y := e.cell.Y - vector[1]

    i, err := resolveXY(X, Y)
    if err != nil {
        // Идём дальше
        return 1
    }

    switch world[i].(type) {
    case *EatCell:
        e.health += 10
        e.score += config.RatingEat
        world[i] = CreateEmptyCell(X, Y)

        //if e.name == "alpha" {
        //    fmt.Printf("Eat\n", )
        //}
    case *PoisonCell:
        e.Die()
        world[i] = CreateEmptyCell(X, Y)
    }

    return 1
}

func (e *LiveCell) Recycle(vector [2]int) int {
    X := e.cell.X + vector[0]
    Y := e.cell.Y - vector[1]

    i, err := resolveXY(X, Y)
    if err != nil {
        // Идём дальше
        return 1
    }

    switch world[i].(type) {
    case *PoisonCell:
        e.score += config.RatingRecycle
        world[i] = CreateEatCell(X, Y)

        //if e.name == "alpha" {
        //    fmt.Printf("Recycle\n")
        //}
    }

    return 1
}

func (e *LiveCell) Attack(vector [2]int) {
    X := e.cell.X + vector[0]
    Y := e.cell.Y - vector[1]

    i, err := resolveXY(X, Y)
    if err != nil {
        // Идём дальше
        return
    }

    switch t := world[i].(type) {
    case *LiveCell:
        t.health -= 10
        if t.health < 1 {
            world[i] = CreateEatCell(X, Y)
            e.score += config.RatingEat
        }
    case *EatCell:
        world[i] = CreateEmptyCell(X, Y)
    case *PoisonCell:
        world[i] = CreateEmptyCell(X, Y)
    case *WellCell:
        e.health -= 10
    }
}

func (e *LiveCell) Die() {
    iod, _ := resolveXY(e.cell.X, e.cell.Y)
    world[iod] = CreateEmptyCell(e.cell.X, e.cell.Y)

    log(fmt.Sprintf("S D %d,%d\n", e.cell.X, e.cell.Y))

}

func (e *LiveCell) Move(vector [2]int) int {

    // movieX, movieY - координаты движения
    // i - адрес в массиве
    movieX := e.cell.X + vector[0]
    movieY := e.cell.Y - vector[1]
    i, err := resolveXY(movieX, movieY)

    // Движение за границы
    if err != nil {
        return 1
    }

    switch world[i].(type) {
    case *EmptyCell:
        // Наткнулись на пустую клетку
        if *PrintActionPtr && *PrintActionLevelPtr > 2 {
            fmt.Printf("[%s.%d] move (%d,%d)\n", e.name, e.health, movieX, movieY)
        }

        log(fmt.Sprintf("S M %d,%d %d,%d\n", e.cell.X, e.cell.Y, movieX, movieY))

        iod, _ := resolveXY(e.cell.X, e.cell.Y)
        world[iod] = CreateEmptyCell(e.cell.X, e.cell.Y)

        e.cell.X = movieX
        e.cell.Y = movieY
        world[i] = e

        //if e.name == "alpha" {
        //    fmt.Printf("Move\n", )
        //}

        e.score += config.RatingMove

    case *PoisonCell:
        // Наступили на яд
        if *PrintActionPtr && *PrintActionLevelPtr > 2 {
            fmt.Printf("[%s.%d] move and die (poison)\n", e.name, e.health)
        }

        e.Die()

    case *EatCell:
        iod, _ := resolveXY(e.cell.X, e.cell.Y)
        world[iod] = CreateEmptyCell(e.cell.X, e.cell.Y)

        // Наступили на еду
        e.cell.X = movieX
        e.cell.Y = movieY
        world[i] = e

        e.score += config.RatingMove

        log(fmt.Sprintf("S M %d,%d %d,%d\n", e.cell.X, e.cell.Y, movieX, movieY))
    }

    // ОК
    return 1
}

func GWaitGenomeGenerator() Genome {
    var genome Genome
    for index := range genome {
        genome[index] = GWait
    }
    return genome
}

func RandGenomeGenerator() Genome {
    var genome Genome
    for index := range genome {
        genome[index] = RandomGen() // пока что геном заполняется рандомно
    }
    return genome
}

// Команды генома:
const (
    GWait = iota

    // Движение
    GMoveUp
    GMoveUpLeft
    GMoveUpRight
    GMoveLeft
    GMoveRight
    GMoveDown
    GMoveDownLeft
    GMoveDownRight

    // Посмотреть
    GSeeUp
    GSeeUpLeft
    GSeeUpRight
    GSeeLeft
    GSeeRight
    GSeeDown
    GSeeDownLeft
    GSeeDownRight

    GEatUp
    GEatUpLeft
    GEatUpRight
    GEatLeft
    GEatRight
    GEatDown
    GEatDownLeft
    GEatDownRight

    //GAttackUp
    //GAttackUpLeft
    //GAttackUpRight
    //GAttackLeft
    //GAttackRight
    //GAttackDown
    //GAttackDownLeft
    //GAttackDownRight

    GRecycleUp
    GRecycleUpLeft
    GRecycleUpRight
    GRecycleLeft
    GRecycleRight
    GRecycleDown
    GRecycleDownLeft
    GRecycleDownRight

    // Конец команд
    GEnd

    GJumpStart
    GJumpEnd = GJumpStart + GenomeSize - 1
)

func (e *LiveCell) Action() bool {
    e.health--

    if e.name == "alpha" {
        fmt.Printf("")
    }

    if !e.IsLive() {
        e.Die()
        return true
    }

    // Выполняем команду в геноме
    switch e.genome[e.it] {

    case GWait:
        // Ничего не делать
        if *PrintActionPtr && *PrintActionLevelPtr > 3 {
            fmt.Printf("[%s.%d] wait\n", e.name, e.health)
        }

    case GSeeUp:
        e.it += e.See(Up)
    case GSeeUpRight:
        e.it += e.See(UpRight)
    case GSeeRight:
        e.it += e.See(Right)
    case GSeeDownRight:
        e.it += e.See(DownRight)
    case GSeeDown:
        e.it += e.See(Down)
    case GSeeDownLeft:
        e.it += e.See(DownLeft)
    case GSeeLeft:
        e.it += e.See(Left)
    case GSeeUpLeft:
        e.it += e.See(UpLeft)

    case GMoveUp:
        e.it += e.Move(Up)
    case GMoveUpRight:
        e.it += e.Move(UpRight)
    case GMoveRight:
        e.it += e.Move(Right)
    case GMoveDownRight:
        e.it += e.Move(DownRight)
    case GMoveDown:
        e.it += e.Move(Down)
    case GMoveDownLeft:
        e.it += e.Move(DownLeft)
    case GMoveLeft:
        e.it += e.Move(Left)
    case GMoveUpLeft:
        e.it += e.Move(UpLeft)

    //case GAttackUp:
    //    e.Attack(Up)
    //case GAttackUpRight:
    //    e.Attack(UpRight)
    //case GAttackRight:
    //    e.Attack(Right)
    //case GAttackDownRight:
    //    e.Attack(DownRight)
    //case GAttackDown:
    //    e.Attack(Down)
    //case GAttackDownLeft:
    //    e.Attack(DownLeft)
    //case GAttackLeft:
    //    e.Attack(Left)
    //case GAttackUpLeft:
    //    e.Attack(UpLeft)

    case GEatUp:
        e.it += e.Eat(Up)
    case GEatUpRight:
        e.it += e.Eat(UpRight)
    case GEatRight:
        e.it += e.Eat(Right)
    case GEatDownRight:
        e.it += e.Eat(DownRight)
    case GEatDown:
        e.it += e.Eat(Down)
    case GEatDownLeft:
        e.it += e.Eat(DownLeft)
    case GEatLeft:
        e.it += e.Eat(Left)
    case GEatUpLeft:
        e.it += e.Eat(UpLeft)

    case GRecycleUp:
        e.it += e.Recycle(Up)
    case GRecycleUpRight:
        e.it += e.Recycle(UpRight)
    case GRecycleRight:
        e.it += e.Recycle(Right)
    case GRecycleDownRight:
        e.it += e.Recycle(DownRight)
    case GRecycleDown:
        e.it += e.Recycle(Down)
    case GRecycleDownLeft:
        e.it += e.Recycle(DownLeft)
    case GRecycleLeft:
        e.it += e.Recycle(Left)
    case GRecycleUpLeft:
        e.it += e.Recycle(UpLeft)

    default:
        if *JumpEnable {
            switch {
            // Это номер в геноме куда переместить указатель
            case e.genome[e.it] >= GJumpStart && e.genome[e.it] <= GJumpEnd:
                // Безусловный переход
                if *PrintActionPtr && *PrintActionLevelPtr > 2 {
                    fmt.Printf("[%s.%d] jumpTo %d\n", e.name, e.health, e.genome[e.it])
                }
                e.it = int(e.genome[e.it]) - GEnd - 1

                //if e.name == "alpha" {
                //    fmt.Printf("Jump to %d\n", e.it)
                //}
            default:
                // Неизвестная команда
                e.it++
            }
        } else {
            // Неизвестная команда
            e.it++
        }
    }

    if  e.it >= GenomeSize {
        e.it = 0
    }

    return !e.IsLive()
}

func (e *LiveCell) PreviewGenome() {
    log(fmt.Sprintf("GENOM %s %d\n", e.name, e.genome))
}
