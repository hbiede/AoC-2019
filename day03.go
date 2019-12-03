package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	partA       = flag.Bool("partA", true, "Perform part A solution?")
	inputFile   = flag.String("inputFile", "inputs/day03.txt", "Input File")
)

type coordinate struct {
	x int
	y int
}

func main() {
	flag.Parse()

	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	//noinspection GoUnhandledErrorResult
	defer file.Close()

	// Read in inputs
	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		log.Fatal("Include 2 lines of wire components")
	}
	wire1 := scanner.Text()
	if !scanner.Scan() {
		log.Fatal("Include 2 lines of wire components")
	}
	wire2 := scanner.Text()

	wire1Components := strings.Split(wire1, ",")
	wire2Components := strings.Split(wire2, ",")

	board := make(map[coordinate]string)
	if *partA {
		// part A
		addWireToBoard(wire1Components, "A", board)
		addWireToBoard(wire2Components, "B", board)

		minDistance := math.MaxInt64
		minPosition := coordinate{0, 0}
		for key, value := range board {
			if value == "X" && distanceFromOrigin(key) < minDistance {
				minPosition = key
				minDistance = distanceFromOrigin(key)
			}
		}
		fmt.Printf("%d spaces from origin (at (%d, %d))", minDistance, minPosition.x, minPosition.y)
	} else {
		// part B
		addWireToBoard(wire1Components, "A", board)
		closestIntersection, combinedSteps := addWireToBoard(wire2Components, "B", board)
		fmt.Printf("%d combined spaces from origin (at (%d, %d))", combinedSteps, closestIntersection.x, closestIntersection.y)
	}
}

func addWireToBoard(components []string, wireName string, board map[coordinate]string) (coordinate, int) {
	coord := coordinate{}
	distance := 0
	closestIntersection := coordinate{}
	closestIntersectionDistanceCombined := math.MaxInt64
	for _, component := range components {
		changeInX := 0
		changeInY := 0
		prefix := ""

		// what direction of travel
		if strings.Contains(component, "U") {
			changeInY = 1
			prefix = "U"
		} else if strings.Contains(component, "D") {
			changeInY = -1
			prefix = "D"
		} else if strings.Contains(component, "R") {
			changeInX = 1
			prefix = "R"
		} else if strings.Contains(component, "L") {
			changeInX = -1
			prefix = "L"
		}

		length, err := strconv.Atoi(strings.TrimPrefix(component, prefix))
		if err != nil {
			log.Fatalf("%s contains no prefix", component)
		}

		for i := 0; i < length; i++ {
			coord.x += changeInX
			coord.y += changeInY
			distance++
			if !strings.Contains(board[coord], wireName) && board[coord] != "" {
				// Already crossed by the other wire, add an intersection
				distanceOnOtherWire, err := strconv.Atoi(strings.Split(board[coord], " - ")[1])
				if err != nil {
					log.Fatal("Don't give bad wire names")
				}
				if distanceOnOtherWire + distance < closestIntersectionDistanceCombined {
					// new closet coordinate by steps
					closestIntersectionDistanceCombined = distanceOnOtherWire + distance
					closestIntersection = coord
				}
				board[coord] = "X"
			} else if !strings.Contains(board[coord], wireName) {
				board[coord] = wireName + " - " + strconv.Itoa(distance)
			}
		}
	}
	return closestIntersection, closestIntersectionDistanceCombined
}

func distanceFromOrigin(coord coordinate) int {
	return int(math.Abs(float64(coord.x)) + math.Abs(float64(coord.y)))
}
