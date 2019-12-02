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

	fmt.Printf("commands[0] = %d\n", IntOpCodeComputer(commands))
}
