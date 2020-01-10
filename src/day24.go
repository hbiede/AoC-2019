package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
)

type board struct {
	bugs [][]bool
}

func (b *board) sim() int {
	oldBugBoard := b.clone()
	for i, row := range b.bugs {
		for j, bugHere := range row {
			adjacentBugs := oldBugBoard.adjacentBugs(i, j)
			if bugHere && adjacentBugs != 1 {
				b.bugs[i][j] = false
			} else if !bugHere && (adjacentBugs == 1 || adjacentBugs == 2) {
				b.bugs[i][j] = true
			}
		}
	}
	return b.diversity()
}

func (b *board) clone() board {
	bugs := make([][]bool, 0)
	for i, row := range b.bugs {
		bugs = append(bugs, make([]bool, len(row)))
		for j, bug := range row {
			bugs[i][j] = bug
		}
	}
	return board{bugs: bugs}
}

func (b *board) adjacentBugs(i int, j int) int {
	if i < 0 || j < 0 || i >= len(b.bugs) || j >= len(b.bugs[i]) {
		return -1
	}

	bugs := 0

	if i - 1 >= 0 && b.bugs[i - 1][j] {
		bugs++
	}
	if j - 1 >= 0 && b.bugs[i][j - 1] {
		bugs++
	}
	if i + 1 < len(b.bugs) && b.bugs[i + 1][j] {
		bugs++
	}
	if j + 1 < len(b.bugs[i]) && b.bugs[i][j + 1] {
		bugs++
	}
	return bugs
}

func (b *board) adjacentBugsRecursive(i int, j int) int {
	bugs := 0

	if i - 1 >= 0 && b.bugs[i - 1][j] {
		bugs++
	}
	if j - 1 >= 0 && b.bugs[i][j - 1] {
		bugs++
	}
	if i + 1 < len(b.bugs) && b.bugs[i + 1][j] {
		bugs++
	}
	if j + 1 < len(b.bugs[i]) && b.bugs[i][j + 1] {
		bugs++
	}
	return bugs
}

func (b *board) diversity() int {
	diversityRating := 0
	for i, row := range b.bugs {
		for j, bugHere := range row {
			if bugHere {
				diversityRating += int(math.Pow(2, float64(i * len(row) + j)))
			}
		}
	}
	return diversityRating
}

var (
	inputFile   = flag.String("inputFile", "inputs/day24.txt", "Input File")
	bugChar = flag.String("bugChar", "#", "The char representing the bug in the input")
)

func main() {
	flag.Parse()

	board := processInput()
	part1(board)
}

func processInput() board {
	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	//noinspection GoUnhandledErrorResult
	defer file.Close()

	// Read in inputs
	scanner := bufio.NewScanner(file)
	bugs := make([][]bool, 0)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]bool, len(line))
		for i, char := range line {
			row[i] = char == []rune(*bugChar)[0]
		}
		bugs = append(bugs, row)
	}
	return board{bugs: bugs}
}

func part1(board board) {
	diversityCountMaps := make(map[int]bool)
	diversityCountMaps[board.diversity()] = true

	for {
		if hasSeen, keyExists := diversityCountMaps[board.sim()]; hasSeen && keyExists {
			fmt.Printf("Duplicate diversity rating of %d\n", board.diversity())
			break
		} else {
			diversityCountMaps[board.diversity()] = true
		}
	}
}
