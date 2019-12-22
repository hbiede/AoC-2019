package main

import (
	"IntCode"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coord struct {
    X int
    Y int
}

var (
    inputFile = flag.String("inputFile", "inputs/day19.txt", "Input File")
    squareWidth = flag.Int("squareWidth", 100, "The width of the square to be saught")
)

func main() {
    flag.Parse()

    commands := processInput()
    runA(commands)
    runB(commands)
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

func runA(commands []int) {
	beamMap := make(map[coord]bool)
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			if check(x, y, commands) {
				beamMap[coord{X: x, Y: y}] = true
			}
		}
	}

	fmt.Printf("%d coords affected (166 expected)\n", len(beamMap))
}

func runB(commands []int) {
    for y := *squareWidth; ; y++ {
    	beamFoundOnX := false
    	for x := 0; ; x++ {
    		if check(x, y, commands) {
			    beamFoundOnX = true
    			if check(x + *squareWidth - 1, y, commands) && check(x, y + *squareWidth - 1, commands) {
				    fmt.Printf("Square found at (%d, %d)... %d (Expected 3790981)\n", x, y, x * 10000 + y)
				    return
			    }
		    } else if beamFoundOnX {
		    	// passed through the beam
		    	break
		    } else {
		    	// still looking for beam
		    	x++
		    }
	    }
    }
}

func check(x int, y int, commands []int) bool {
	botController := IntCode.New()
	commandsClone := make([]int, len(commands))
	copy(commandsClone, commands)
	go botController.Run(commandsClone)
	botController.Input <- x
	botController.Input <- y
	return 1 == <- botController.Output
}
