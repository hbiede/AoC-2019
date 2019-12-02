package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const DesiredOutput = 19690720

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read in input
	scanner := bufio.NewScanner(file)
	inputString := ""
	for scanner.Scan() {
		inputString += scanner.Text()
	}

	commandStrings := strings.Split(inputString, ",")
	var commands []int
	for _, commandString := range commandStrings { // the _ disregards the index and keeps the element in commandString
		command, err := strconv.Atoi(commandString)
		if err != nil {
			log.Fatal(err)
		}
		commands = append(commands, command)
	}

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			commandsClone := make([]int, len(commands))
			copy(commandsClone, commands)
			commandsClone[1] = i
			commandsClone[2] = j
			if intOpCodeComputer(commandsClone) == DesiredOutput {
				fmt.Printf("Noun: %d\tVerb: %d. 100 * noun + verb: %d", i, j, i * 100 + j)
			}
		}
	}
}

/**
 * Returns the value in commands[0] after program completion, -1 if 99 is never reached
 */
func intOpCodeComputer(commands []int) int {
	// Process program
	for i := 0; i + 3 < len(commands); i += 4 {
		switch commands[i] {
		case 1:
			commands[commands[i + 3]] = commands[commands[i + 1]] + commands[commands[i + 2]]
		case 2:
			commands[commands[i + 3]] = commands[commands[i + 1]] * commands[commands[i + 2]]
		case 99:
			return commands[0]
		default:
			log.Fatalf("%d is an unknown command\n", commands[i])
		}
	}
	return -1
}
