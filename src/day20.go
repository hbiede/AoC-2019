package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	inputFile   = flag.String("inputFile", "inputs/day20.txt", "Input File")
)

type coord3d struct {
	x int
	y int
	z int
}

type square struct {
	topLeft     coord3d
	bottomRight coord3d
}

type state struct {
	position coord3d
	distance int
}

type portal struct {
	label string
	from  coord3d
	to    coord3d
	outer bool
}

func getGrid(inputMap []string) (grid map[coord3d]bool, portals map[coord3d]portal, start, end coord3d) {
	outerBorder, innerBorder := square{topLeft: coord3d{2, 2, 0}, bottomRight: coord3d{len(inputMap[0]) - 3, len(inputMap) - 3, 0}}, square{}
	for y := outerBorder.topLeft.y; ; y++ {
		if strings.Contains(inputMap[y][outerBorder.topLeft.x:outerBorder.bottomRight.x], " ") {
			innerBorder.topLeft = coord3d{strings.Index(inputMap[y][outerBorder.topLeft.x:outerBorder.bottomRight.x], " ") + outerBorder.topLeft.x - 1, y - 1, 0}
			innerBorder.bottomRight = coord3d{outerBorder.bottomRight.x - innerBorder.topLeft.x + outerBorder.topLeft.x, outerBorder.bottomRight.y - innerBorder.topLeft.y + outerBorder.topLeft.y, 0}
			break
		}
	}

	grid, portals = make(map[coord3d]bool), make(map[coord3d]portal)
	for y := outerBorder.topLeft.y; y <= outerBorder.bottomRight.y; y++ {
		for x := outerBorder.topLeft.x; x <= outerBorder.bottomRight.x; x++ {
			if inputMap[y][x] == '.' {
				grid[coord3d{x, y, 0}] = true

				var label string
				var pos coord3d
				var outerPortal bool
				if y == outerBorder.topLeft.y {
					label = inputMap[y-2][x:x+1] + inputMap[y-1][x:x+1]
					pos = coord3d{x, y - 1, 0}
					outerPortal = true
				} else if y == outerBorder.bottomRight.y {
					label = inputMap[y+1][x:x+1] + inputMap[y+2][x:x+1]
					pos = coord3d{x, y + 1, 0}
					outerPortal = true
				} else if x == outerBorder.topLeft.x {
					label = inputMap[y][x-2 : x]
					pos = coord3d{x - 1, y, 0}
					outerPortal = true
				} else if x == outerBorder.bottomRight.x {
					label = inputMap[y][x+1 : x+3]
					pos = coord3d{x + 1, y, 0}
					outerPortal = true
				} else if y == innerBorder.bottomRight.y && x > innerBorder.topLeft.x && x < innerBorder.bottomRight.x {
					label = inputMap[y-2][x:x+1] + inputMap[y-1][x:x+1]
					pos = coord3d{x, y - 1, 0}
				} else if y == innerBorder.topLeft.y && x > innerBorder.topLeft.x && x < innerBorder.bottomRight.x {
					label = inputMap[y+1][x:x+1] + inputMap[y+2][x:x+1]
					pos = coord3d{x, y + 1, 0}
				} else if x == innerBorder.bottomRight.x && y > innerBorder.topLeft.y && y < innerBorder.bottomRight.y {
					label = inputMap[y][x-2 : x]
					pos = coord3d{x - 1, y, 0}
				} else if x == innerBorder.topLeft.x && y > innerBorder.topLeft.y && y < innerBorder.bottomRight.y {
					label = inputMap[y][x+1 : x+3]
					pos = coord3d{x + 1, y, 0}
				}

				// add labels for portals
				if label == "AA" {
					start = coord3d{x, y, 0}
				} else if label == "ZZ" {
					end = coord3d{x, y, 0}
				} else if label != "" {
					portals[pos] = portal{label: label, from: coord3d{x, y, 0}, outer: outerPortal}
					grid[pos] = true
				}
			}
		}
	}

	for i, p1 := range portals {
		for _, p2 := range portals {
			if p1.label == p2.label && p1.from != p2.from {
				p1.to = p2.from
				portals[i] = p1
			}
		}
	}

	return
}

func part1(inputMap []string) {
	grid, portals, start, end := getGrid(inputMap)

	directions := []coord3d{{0, -1, 0}, {1, 0, 0}, {0, 1, 0}, {-1, 0, 0}}
	queue, visited := []state{state{position: start}}, map[coord3d]bool{start: true}
	var currState state
	for {
		currState, queue = queue[0], queue[1:]
		for _, d := range directions {
			next := coord3d{currState.position.x + d.x, currState.position.y + d.y, 0}

			if next == end {
				fmt.Printf("Steps to ZZ: %d (Expected 476)", currState.distance + 1)
				return
			}

			if grid[next] && !visited[next] {
				visited[next] = true

				p, ok := portals[next]
				if ok {
					next = p.to
				}

				queue = append(queue, state{next, currState.distance + 1})
			}
		}
	}
}

func part2(inputMap []string) {
	grid, portals, start, end := getGrid(inputMap)

	directions := []coord3d{{0, -1, 0}, {1, 0, 0}, {0, 1, 0}, {-1, 0, 0}}
	queue, visited := []state{state{position: start}}, map[coord3d]bool{coord3d{start.x, start.y, 0}: true}
	var st state
	for {
		st, queue = queue[0], queue[1:]
		for _, d := range directions {
			next := coord3d{st.position.x + d.x, st.position.y + d.y, st.position.z}

			if next == end {
				fmt.Printf("Steps to recursive exit: %d (Expected 5350)", st.distance + 1)
				return
			}

			if grid[coord3d{next.x, next.y, 0}] && !visited[next] {
				visited[next] = true

				p, ok := portals[coord3d{next.x, next.y, 0}]
				if ok && (st.position.z > 0 || !p.outer) {
					next = coord3d{p.to.x, p.to.y, st.position.z}
					if p.outer {
						next.z--
					} else {
						next.z++
					}

					visited[next] = true
				}

				queue = append(queue, state{next, st.distance + 1})
			}
		}
	}
}

func main() {
	inputMap := processInput()
	part1(inputMap)
	part2(inputMap)
}

func processInput() []string {
	flag.Parse()
	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	//noinspection GoUnhandledErrorResult
	defer file.Close()

	// Read in inputs
	scanner := bufio.NewScanner(file)
	inputMap := make([]string, 0)
	for scanner.Scan() {
		inputMap = append(inputMap, scanner.Text())
	}

	return inputMap
}