package main

import (
    "IntCode"
    "bufio"
    "flag"
    "fmt"
    "image"
    "image/color"
    "image/draw"
    "image/png"
    "log"
    "math"
    "os"
    "strconv"
    "strings"
    "sync"
)

const (
    UP    = 0
    RIGHT = 1
    DOWN  = 2
    LEFT  = 3

    COUNTERCLOCKWISE = 0
    CLOCKWISE        = 1
)

type coord struct {
    X int
    Y int
}

func (c *coord) clone() coord {
    return coord{X: c.X, Y: c.Y}
}

func (c *coord) move(direction int) {
    switch direction {
    case UP:
        c.Y++
    case RIGHT:
        c.X++
    case DOWN:
        c.Y--
    case LEFT:
        c.X--
    default:
        log.Fatalf("Invalid direction received: %d", direction)
    }
}

var (
    partA           = flag.Bool("partA", false, "Is Part A?")
    inputFile       = flag.String("inputFile", "inputs/day11.txt", "Input File")
    blackInt        = flag.Int("blackInt", 0, "The Color Black")
    whiteInt        = flag.Int("whiteInt", 1, "The Color Black")
    defaultColorInt = flag.Int("defaultColorInt", *blackInt, "The Color White")
)

func main() {
    flag.Parse()

    commands := processInput()

    var outputLoc, expectation string
    if *partA {
        outputLoc = "outputs/day11-1.png"
        expectation = "Gibberish expected on 2211 spots"
    } else {
        outputLoc = "outputs/day11-2.png"
        expectation = "EFCKUEGC expected"
    }

    f, err := os.Create(outputLoc)
    if err != nil {
        log.Fatalln(err)
    }
    err = png.Encode(f, simulateAnGenerateImage(commands))
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println("Image Generated. ", expectation)
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

func simulateAnGenerateImage(commands []int) draw.Image {
    robot := IntCode.New()
    commandsClone := make([]int, len(commands))
    copy(commandsClone, commands)
    go robot.Run(commandsClone)

    robotPaintJob := make(map[coord]int)
    robotFacingDirection := UP
    robotLocation := coord{X: 0, Y: 0}
    if !*partA {
        robotPaintJob[robotLocation] = *whiteInt
    }
    robot.Input <- robotPaintJob[robotLocation]

    wg := sync.WaitGroup{}
    wg.Add(1)
    go func(robot *IntCode.Stream) {
        defer func() {
            wg.Done()
            recovery := recover()
            if recovery != nil && fmt.Sprint(recovery) != "send on closed channel" {
                panic(recovery)
            }
        }()

        for paintColor := range robot.Output {
            robotPaintJob[robotLocation.clone()] = paintColor

            turnDirection := <-robot.Output
            if turnDirection == COUNTERCLOCKWISE {
                robotFacingDirection -= 1
                if robotFacingDirection == -1 {
                    robotFacingDirection = LEFT
                }
            } else if turnDirection == CLOCKWISE {
                robotFacingDirection += 1
                if robotFacingDirection == 4 {
                    robotFacingDirection = UP
                }
            } else {
                log.Fatalf("Invalid turn direction received: %d", turnDirection)
            }

            robotLocation.move(robotFacingDirection)

            if _, keyExists := robotPaintJob[robotLocation]; keyExists {
                robot.Input <- robotPaintJob[robotLocation]
            } else {
                robot.Input <- *defaultColorInt
            }
        }
    }(robot)
    wg.Wait()

    fmt.Printf("%d spots painted\n", len(robotPaintJob))

    return createImage(robotPaintJob)
}

func createImage(coords map[coord]int) draw.Image {
    minCoord, maxCoord := findMinAndMax(coords)

    topLeft := image.Point{X: 0, Y: 0}
    bottomRight := image.Point{X: maxCoord.X - minCoord.X + 1, Y: maxCoord.Y - minCoord.Y + 1}
    img := image.NewRGBA(image.Rectangle{Min: topLeft, Max: bottomRight})
    whiteColor := color.White
    blackColor := color.Black
    var defaultColor color.Color
    if *defaultColorInt == *blackInt {
        defaultColor = blackColor
    } else {
        defaultColor = whiteColor
    }

    for i := 0; i <= maxCoord.X-minCoord.X; i++ {
        for j := 0; j <= maxCoord.Y-minCoord.Y; j++ {
            img.Set(i, j, defaultColor)
        }
    }

    for coordinate, colorInt := range coords {
        var setColor color.Color
        if colorInt == *blackInt {
            setColor = blackColor
        } else {
            setColor = whiteColor
        }
        img.Set(coordinate.X-minCoord.X, maxCoord.Y-coordinate.Y, setColor)
    }
    return img
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
