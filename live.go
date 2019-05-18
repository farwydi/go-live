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
        return 2
    case *PoisonCell:
        return 3
    case *EmptyCell:
        return 4
    case *WellCell:
        return 5
    }

    return 1
}

func (e *LiveCell) Eat(vector [2]int) {
    X := e.cell.X + vector[0]
    Y := e.cell.Y - vector[1]

    i, err := resolveXY(X, Y)
    if err != nil {
        // Идём дальше
        return
    }

    switch world[i].(type) {
    case *EatCell:
        e.health += 10
    case *PoisonCell:
        e.health = 0
    case *EmptyCell:
        e.health -= 1
    case *WellCell:
        e.health -= 10
    }
}

func (e *LiveCell) Recycle(vector [2]int) {
    X := e.cell.X + vector[0]
    Y := e.cell.Y - vector[1]

    i, err := resolveXY(X, Y)
    if err != nil {
        // Идём дальше
        return
    }

    switch world[i].(type) {
    case *EatCell:
        e.health -= 10
        world[i] = CreatePoisonCell(X, Y)
    case *PoisonCell:
        world[i] = CreateEatCell(X, Y)
    case *EmptyCell:
        e.health -= 1
    case *WellCell:
        e.health -= 10
    }
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
    case *EatCell:
        world[i] = CreateEmptyCell(X, Y)
    case *PoisonCell:
        world[i] = CreateEmptyCell(X, Y)
    case *EmptyCell:
        e.health -= 1
    case *WellCell:
        e.health -= 10
    }
}

func (e *LiveCell) Move(vector [2]int) error {

    // movieX, movieY - координаты движения
    // i - адрес в массиве
    movieX := e.cell.X + vector[0]
    movieY := e.cell.Y - vector[1]
    i, err := resolveXY(movieX, movieY)

    // Движение за границы
    if err != nil {
        return err
    }

    switch t := world[i].(type) {
    case *EmptyCell:
        // Наткнулись на пустую клетку
        if *PrintActionPtr && *PrintActionLevelPtr > 2 {
            fmt.Printf("[%s.%d] move (%d,%d)\n", e.name, e.health, movieX, movieY)
        }

        log(fmt.Sprintf("S M %d,%d %d,%d\n", e.cell.X, e.cell.Y, movieX, movieY))

        mutex.Lock()
        e.cell.X = movieX
        e.cell.Y = movieY
        mutex.Unlock()

        e.score += config.RatingMove

    case *PoisonCell:
        // Наступили на яд
        if *PrintActionPtr && *PrintActionLevelPtr > 2 {
            fmt.Printf("[%s.%d] move and die (poison)\n", e.name, e.health)
        }

        e.health = 0

        log(fmt.Sprintf("S D %d,%d\n", e.cell.X, e.cell.Y))

    case *EatCell:
        // Наступили на еду
        if *PrintActionPtr && *PrintActionLevelPtr > 2 {
            fmt.Printf("[%s.%d] move and eat\n", e.name, e.health)
        }

        // Забираем калории еды и добавляем их к текущему ХП
        e.health += t.calories
        // Начисляем рейтинг за активность
        e.score += config.RatingEat
        // Обнуляет клетку с едой
        world[i] = CreateEmptyCell(movieX, movieY)
        // Переходим на эту клетку

        // Двигаем клетку на место с едой
        log(fmt.Sprintf("S E %d,%d %d,%d\n", e.cell.X, e.cell.Y, movieX, movieY))

        mutex.Lock()
        e.cell.X = movieX
        e.cell.Y = movieY
        mutex.Unlock()
    }

    // ОК
    return nil
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

    GAttackUp
    GAttackUpLeft
    GAttackUpRight
    GAttackLeft
    GAttackRight
    GAttackDown
    GAttackDownLeft
    GAttackDownRight

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
    GJumpEnd = GJumpStart + (GEnd - 1)
)

func (e *LiveCell) Action() bool {
    //log(fmt.Sprintf("GENOM %s %d\n", e.name, e.genome))

    e.health--

    if !e.IsLive() {
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
        e.Move(Up)
    case GMoveUpRight:
        e.Move(UpRight)
    case GMoveRight:
        e.Move(Right)
    case GMoveDownRight:
        e.Move(DownRight)
    case GMoveDown:
        e.Move(Down)
    case GMoveDownLeft:
        e.Move(DownLeft)
    case GMoveLeft:
        e.Move(Left)
    case GMoveUpLeft:
        e.Move(UpLeft)

    case GAttackUp:
        e.Attack(Up)
    case GAttackUpRight:
        e.Attack(UpRight)
    case GAttackRight:
        e.Attack(Right)
    case GAttackDownRight:
        e.Attack(DownRight)
    case GAttackDown:
        e.Attack(Down)
    case GAttackDownLeft:
        e.Attack(DownLeft)
    case GAttackLeft:
        e.Attack(Left)
    case GAttackUpLeft:
        e.Attack(UpLeft)

    case GEatUp:
        e.Eat(Up)
    case GEatUpRight:
        e.Eat(UpRight)
    case GEatRight:
        e.Eat(Right)
    case GEatDownRight:
        e.Eat(DownRight)
    case GEatDown:
        e.Eat(Down)
    case GEatDownLeft:
        e.Eat(DownLeft)
    case GEatLeft:
        e.Eat(Left)
    case GEatUpLeft:
        e.Eat(UpLeft)

    case GRecycleUp:
        e.Recycle(Up)
    case GRecycleUpRight:
        e.Recycle(UpRight)
    case GRecycleRight:
        e.Recycle(Right)
    case GRecycleDownRight:
        e.Recycle(DownRight)
    case GRecycleDown:
        e.Recycle(Down)
    case GRecycleDownLeft:
        e.Recycle(DownLeft)
    case GRecycleLeft:
        e.Recycle(Left)
    case GRecycleUpLeft:
        e.Recycle(UpLeft)

    default:
        switch {
        // Это номер в геноме куда переместить указатель
        case e.it >= GJumpStart && e.it <= GJumpEnd:
            // Безусловный переход
            if *PrintActionPtr && *PrintActionLevelPtr > 2 {
                fmt.Printf("[%s.%d] jumpTo %d\n", e.name, e.health, e.it)
            }
            e.it -= GEnd + 1
        default:
            // Неизвестная команда
            e.it++
        }
    }

    return false
}
