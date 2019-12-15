package main

import (
	"IntCode"
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
)

const (
	NORTH = 1
	SOUTH = 2
	WEST  = 3
	EAST  = 4
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
		c.Y++
	case EAST:
		c.X++
	case SOUTH:
		c.Y--
	case WEST:
		c.X--
	default:
		log.Fatalf("Invalid direction received: %d", direction)
	}
	return *c
}

func (c *coord) closest(movementMap map[coord]int) int {
	minSteps := math.MaxInt64
	if testDist := movementMap[c.clone().move(NORTH)]; testDist < minSteps && testDist > 0 {
		minSteps = testDist
	}
	if testDist := movementMap[c.clone().move(SOUTH)]; testDist < minSteps && testDist > 0 {
		minSteps = testDist
	}
	if testDist := movementMap[c.clone().move(EAST)]; testDist < minSteps && testDist > 0 {
		minSteps = testDist
	}
	if testDist := movementMap[c.clone().move(WEST)]; testDist < minSteps && testDist > 0 {
		minSteps = testDist
	}
	return minSteps
}

func (c *coord) hasAdjacentUnsolved(movementMap map[coord]int) bool {
	for i := 1; i <= 4; i++ {
		if _, keyExists := movementMap[c.clone().move(i)]; !keyExists {
			return true
		}
	}
	return false
}

var (
	inputFile = flag.String("inputFile", "inputs/day15.txt", "Input File")
)

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

	movementMap := make(map[coord]int)
	moveDirection := NORTH
	botController.Input <- moveDirection
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

		robotLocation := coord{X: 0, Y: 0}
		movementMap[robotLocation] = 1
		for movementStatus := range botController.Output {
			if movementStatus != 0 {
				robotLocation = robotLocation.move(moveDirection)
				stepsToLocation := robotLocation.closest(movementMap) + 1
				if _, keyExists := movementMap[robotLocation]; !keyExists || stepsToLocation < movementMap[robotLocation] && stepsToLocation > 0 {
					movementMap[robotLocation] = stepsToLocation
				}

				if movementStatus == 2 {
					fmt.Printf("Movements to the oxygen port (262 expected): %d\n", movementMap[robotLocation] - 1)
					break
				}
			}
			moveDirection = rand.Intn(4) + 1
			botController.Input <- moveDirection
		}

		fmt.Println("Warning: This is a poor strategy that utilizes randomness to find all branches of the maze")

		moveDirection = rand.Intn(4) + 1
		botController.Input <- moveDirection
		movements := 0
		movementMap = make(map[coord]int)
		movementMap[robotLocation] = 1
		for movementStatus := range botController.Output {
			movements++
			if movementStatus != 0 {
				robotLocation = robotLocation.move(moveDirection)
				stepsToLocation := robotLocation.closest(movementMap) + 1
				if _, keyExists := movementMap[robotLocation]; !keyExists || stepsToLocation < movementMap[robotLocation] && stepsToLocation > 0 {
					movementMap[robotLocation] = stepsToLocation
				} else if keyExists {
				}

				if movements % 1000 == 0 {
					tileRemainsUnexplored := false
					maxDistance := math.MinInt64
					for coordinate, distance := range movementMap {
						tileRemainsUnexplored = (coordinate.hasAdjacentUnsolved(movementMap) && distance > -1) || tileRemainsUnexplored
						if distance > maxDistance && distance > 0 {
							maxDistance = distance
						}
					}

					if !tileRemainsUnexplored {
						fmt.Printf("Minutes to fill the chamber (expected 314): %d\n", maxDistance - 1)
						break
					}
				}
			} else {
				if _, keyExists := movementMap[robotLocation.clone().move(moveDirection)]; !keyExists {
					movementMap[robotLocation.clone().move(moveDirection)] = -1
				}
			}
			moveDirection = rand.Intn(4) + 1
			botController.Input <- moveDirection
		}
	}(botController)
	wg.Wait()
}
