package main

import (
    "IntCode"
    "bufio"
    "flag"
    "fmt"
    "log"
    "math"
    "os"
    "strconv"
    "strings"
    "sync"
)

var (
    partA     = flag.Bool("partA", true, "Perform part A solution?")
    inputFile = flag.String("inputFile", "inputs/day13.txt", "Input File")
)

type coord struct {
    X int
    Y int
}

func (c *coord) clone() coord {
    return coord{X: c.X, Y: c.Y}
}

func main() {
    flag.Parse()

    commands := processInput()
    play(commands)
}

func processInput() []int {
    file, err := os.Open(*inputFile)
    if err != nil {
        log.Fatal(err)
    }
    //noinspection GoUnhandledErrorResult
    defer file.Close()

    // Read in inputs
    scanner := bufio.NewScanner(file)
    inputStringFromFile := ""
    for scanner.Scan() {
        inputStringFromFile += scanner.Text()
    }

    commandStrings := strings.Split(inputStringFromFile, ",")
    var commands []int
    for _, commandString := range commandStrings { // the _ disregards the index and keeps the element in commandString
        command, err := strconv.Atoi(commandString)
        if err != nil {
            log.Fatal(err)
        }
        commands = append(commands, command)
    }
    return commands
}

func play(commands []int) {
    gameController := IntCode.New()
    commandsClone := make([]int, len(commands))
    copy(commandsClone, commands)
    if !*partA {
        commandsClone[0] = 2 // insert quarters
    }
    go gameController.Run(commandsClone)

    gameMap := make(map[coord]int)
    wg := sync.WaitGroup{}
    wg.Add(1)
    go func(gameController *IntCode.Stream) {
        defer func() {
            wg.Done()
            recovery := recover()
            if recovery != nil && fmt.Sprint(recovery) != "send on closed channel" {
                panic(recovery)
            }
        }()
        score := 0
        paddleX := 0
        ballX := 0
        for x := range gameController.Output {
            y := <-gameController.Output
            tileID := <-gameController.Output
            if x == -1 && y == 0 {
                score = tileID
            } else {
                gameMap[coord{X: x, Y: y}] = tileID
                if tileID == 3 {
                    paddleX = x
                } else if tileID == 4 {
                    ballX = x
                    if paddleX > ballX {
                        gameController.Input <- -1
                    } else if paddleX < ballX {
                        gameController.Input <- 1
                    } else {
                        gameController.Input <- 0
                    }
                    printGameScreen(createMatrix(gameMap), score)
                }
            }
        }
        if !*partA {
            fmt.Printf("Final score (19210 expected): %d\n", score)
        }
    }(gameController)
    wg.Wait()

    if *partA {
        blockCount := 0
        for _, tileType := range gameMap {
            if tileType == 2 {
                blockCount++
            }
        }
        fmt.Printf("%d blocks on exit (369 expected)", blockCount)
    }
}

func printGameScreen(matrix [][]int, score int) {
    printScreen := fmt.Sprintf("Score: %d\n", score)
    for i := len(matrix) - 1; i >= 0; i-- {
        for _, tile := range matrix[i] {
            switch tile {
            case 0:
                printScreen += " "
            case 1:
                printScreen += "|"
            case 2:
                printScreen += "X"
            case 3:
                printScreen += "–"
            case 4:
                printScreen += "*"
            }
        }
        printScreen += "\n"
    }
    fmt.Println(printScreen)
}

func createMatrix(coords map[coord]int) [][]int {
    minCoord, maxCoord := findMinAndMax(coords)

    matrix := make([][]int, maxCoord.Y-minCoord.Y+1)
    for i := range matrix {
        matrix[i] = make([]int, maxCoord.X-minCoord.X+1)
    }

    for coordinate, tileType := range coords {
        matrix[maxCoord.Y-coordinate.Y][coordinate.X-minCoord.X] = tileType
    }
    return matrix
}

func findMinAndMax(coords map[coord]int) (minCoord coord, maxCoord coord) {
    minCoord = coord{X: math.MaxInt64, Y: math.MaxInt64}
    maxCoord = coord{X: math.MinInt64, Y: math.MinInt64}

    for coordinate := range coords {
        if coordinate.X < minCoord.X {
            minCoord.X = coordinate.X
        }
        if coordinate.Y < minCoord.Y {
            minCoord.Y = coordinate.Y
        }

        if coordinate.X > maxCoord.X {
            maxCoord.X = coordinate.X
        }
        if coordinate.Y > maxCoord.Y {
            maxCoord.Y = coordinate.Y
        }
    }
    return minCoord, maxCoord
}
