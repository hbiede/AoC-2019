package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

	// Process program
	for i := 0; i + 3 < len(commands); i += 4 {
		switch commands[i] {
		case 1:
			commands[commands[i + 3]] = commands[commands[i + 1]] + commands[commands[i + 2]]
		case 2:
			commands[commands[i + 3]] = commands[commands[i + 1]] * commands[commands[i + 2]]
		case 99:
			fmt.Println("Finished Program:")
			fmt.Print(commands[0])
			for j := 1; j < len(commands); j++ {
				fmt.Printf(", %d", commands[j])
			}
		default:
			log.Fatalf("%d is an unknown command\n", commands[i])
		}
	}
}
