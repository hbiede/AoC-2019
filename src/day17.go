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

const (
    NORTH = 0
    EAST  = 1
    SOUTH = 2
    WEST  = 3

    SCAFFOLD = '#'
)

var (
    inputFile = flag.String("inputFile", "inputs/day17.txt", "Input File")
    debug     = flag.Bool("debug", false, "Should I print all the debug info?")
)

type coord struct {
    X int
    Y int
}

func (c *coord) clone() *coord {
    return &coord{X: c.X, Y: c.Y}
}

func (c *coord) move(direction int) coord {
    switch direction {
    case NORTH:
        c.Y-- // Y positive is in the downward direction for this map
    case EAST:
        c.X++
    case SOUTH:
        c.Y++
    case WEST:
        c.X--
    default:
        log.Fatalf("Invalid direction received: %d", direction)
    }
    return *c
}

func (c *coord) canMove(direction int, scaffoldMap map[coord]rune) bool {
    testCoord := c.clone()
    switch direction {
    case NORTH:
        testCoord.Y-- // Y positive is in the downward direction for this map
    case EAST:
        testCoord.X++
    case SOUTH:
        testCoord.Y++
    case WEST:
        testCoord.X--
    default:
        log.Fatalf("Invalid direction received: %d", direction)
    }
    return scaffoldMap[*testCoord] == SCAFFOLD
}

func (c *coord) isIntersection(scaffoldMap map[coord]rune) bool {
    for i := NORTH; i <= WEST; i++ {
        if scaffoldMap[c.clone().move(i)] != SCAFFOLD {
            return false
        }
    }
    return true
}

func (c *coord) findTurnDirection(scaffoldMap map[coord]rune, currentlyFacing int) int {
    right := (currentlyFacing + 1) % 4
    left := currentlyFacing - 1
    for left < 0 {
        left += 4
    }
    aboutFace := (right + 1) % 4
    if c.canMove(right, scaffoldMap) {
        return right
    } else if c.canMove(left, scaffoldMap) {
        return left
    } else if c.canMove(aboutFace, scaffoldMap) {
        return aboutFace
    }
    return currentlyFacing // end of the line
}

func (c *coord) equals(other coord) bool {
    return c.Y == other.Y && c.X == other.X
}

func main() {
    flag.Parse()

    commands := processInput()
    run(commands)
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
    for _, commandString := range commandStrings {
        command, err := strconv.Atoi(commandString)
        if err != nil {
            log.Fatal(err)
        }
        commands = append(commands, command)
    }
    return commands
}

func run(commands []int) {
    botController := IntCode.New()
    commandsClone := make([]int, len(commands))
    copy(commandsClone, commands)
    go botController.Run(commandsClone)

    scaffoldMap := make(map[coord]rune)
    wg := sync.WaitGroup{}
    wg.Add(1)
    go func(botController *IntCode.Stream) {
        defer func() {
            wg.Done()
            recovery := recover()
            if recovery != nil && fmt.Sprint(recovery) != "send on closed channel" {
                panic(recovery)
            }
        }()

        printLocation := coord{X: 0, Y: 0}
        for movementStatus := range botController.Output {
            scaffoldMap[printLocation] = rune(movementStatus)
            fmt.Printf(" %c", scaffoldMap[printLocation])
            if movementStatus == 10 {
                printLocation = coord{X: 0, Y: printLocation.Y + 1}
            } else {
                printLocation = coord{X: printLocation.X + 1, Y: printLocation.Y}
            }
        }
    }(botController)
    wg.Wait()

    alignmentParamSum := 0
    intersectionsFound := 0
    robotLocation := coord{X: -1, Y: -1} // needed for Part 2
    var robotFacing int
    for key, value := range scaffoldMap {
        if value == SCAFFOLD && key.isIntersection(scaffoldMap) {
            alignmentParamSum += key.X * key.Y
            intersectionsFound++
        } else if value == 'v' || value == '^' || value == '<' || value == '>' {
            robotLocation = key
            switch value {
            case 'v':
                robotFacing = SOUTH
            case '^':
                robotFacing = NORTH
            case '<':
                robotFacing = WEST
            default:
                robotFacing = EAST
            }
        }
    }
    fmt.Printf("\nSum of Alignment Parameters is %d from %d intersections (6448 and 12 expected respectively)\n", alignmentParamSum, intersectionsFound)

    // Part 2
    if !robotLocation.equals(coord{X: -1, Y: -1}) {
        movementDirection := robotLocation.findTurnDirection(scaffoldMap, robotFacing)
        commandString := getTurnString(robotFacing, movementDirection)

        pathLeftToTraverse := true
        for pathLeftToTraverse {
            moveLength := 0
            for robotLocation.canMove(movementDirection, scaffoldMap) {
                moveLength++
                robotLocation.move(movementDirection)
            }
            commandString += "," + strconv.Itoa(moveLength)

            newDirection := robotLocation.findTurnDirection(scaffoldMap, movementDirection)
            if newDirection == movementDirection || math.Abs(float64(newDirection-movementDirection)) == 2 {
                pathLeftToTraverse = false // end of the line
            } else {
                commandString += "," + getTurnString(movementDirection, newDirection)
                movementDirection = newDirection
            }
        }
        commandString = strings.Trim(commandString, ",")
        fmt.Printf("Full movement pattern: %s\n", commandString)

        stringIsReducible := true
        nextFunctionName := 'A'
        functions := make(map[rune]string)
        for stringIsReducible {
            stringComponents := make(map[string]int)
            for i := 0; i < len(commandString)-2; i++ {
                for j := i + 1; j < len(commandString)-i && j-i <= 20; j++ {
                    // take the substring of commandString between i and j inclusive
                    // (normally a slice would exclude the second value, hence the "+ 1")
                    // and then increment that string's value in the map. j also starts at i + 1
                    // to prevent any strings of length 1 entering the map
                    substring := string([]rune(commandString)[i : j+1])
                    substringContainsAnotherFunction := false
                    for functionName := range functions {
                        substringContainsAnotherFunction = substringContainsAnotherFunction || strings.Contains(substring, string(functionName))
                    }
                    if !substringContainsAnotherFunction && len(substring) == len(strings.Trim(substring, ",")) {
                        // don't reduce by strings starting or ending with commas
                        stringComponents[substring] = strings.Count(commandString, substring)
                    }
                }
            }

            mostWorthwhileString := ""
            for key, value := range stringComponents {
                // use the length of the string times the number of times the string appears as a "weight"
                // for the string. Short strings are more worth while than long ones to reduce if they appear
                // way more frequently.
                //
                // Uses length minus one because 1 is the length of the replacement and should be accounted for
                // in the replacement arithmetic
                if (len(key)-1)*value > (len(mostWorthwhileString)-1)*stringComponents[mostWorthwhileString] {
                    mostWorthwhileString = key
                }
            }

            if len(mostWorthwhileString) < 2 || len(mostWorthwhileString) == len(commandString) {
                stringIsReducible = false
            } else {
                commandString = strings.ReplaceAll(commandString, mostWorthwhileString, string(nextFunctionName))
                functions[nextFunctionName] = mostWorthwhileString + "\n" // add terminating new line character so IntCode computer can process separate functions
	            if *debug {
		            fmt.Printf("New Function %c (%d replacements): %s\n", nextFunctionName, stringComponents[mostWorthwhileString], mostWorthwhileString)
		            fmt.Printf("New Command String: %s\n\n", commandString)
	            }
                nextFunctionName++
            }
        }
        commandString += "\n" // add terminating new line character so IntCode computer can process separate functions

        botController = IntCode.New()
        copy(commandsClone, commands)
        commandsClone[0] = 2
        go botController.Run(commandsClone)

        wg.Add(1)
        go func(botController *IntCode.Stream) {
            defer func() {
                wg.Done()
                recovery := recover()
                if recovery != nil && fmt.Sprint(recovery) != "send on closed channel" {
                    panic(recovery)
                }
            }()
            for output := range botController.Output {
                if *debug {
                	fmt.Printf("%c", output)
                }
                if output > 127 {
                	fmt.Printf("%d dust cleaned (914900 expected)", output)
                }
            }
        }(botController)

        passFunctionToBot(botController, commandString)

        for _, function := range functions {
            passFunctionToBot(botController, function)
        }

        passFunctionToBot(botController, "n\n") // don't want live feed output

        wg.Wait()
    }
}

func getTurnString(movementDirection int, newDirection int) string {
    if newDirection-movementDirection == 1 || newDirection-movementDirection == NORTH-WEST {
        return "R"
    } else if newDirection-movementDirection == -1 || newDirection-movementDirection == WEST-NORTH {
        return "L"
    } else if newDirection-movementDirection == 2 || newDirection-movementDirection == -2 {
        return "R,R"
    } else if newDirection == movementDirection {
        return ""
    } else {
        log.Fatalf("Unknown turn quantity: %d from %d", newDirection, movementDirection)
        return "" // appease the compiler. Unreachable code
    }
}

func passFunctionToBot(botController *IntCode.Stream, function string) {
    for _, character := range function {
        botController.Input <- int(character)
    }
}
