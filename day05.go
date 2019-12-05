package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"

	"./IntCode"
)

var (
	partA       = flag.Bool("partA", true, "Perform part A solution?")
	inputFile   = flag.String("inputFile", "inputs/day05.txt", "Input File")
)

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

	if *partA {
		// part A
		IntCode.IntOpCodeComputer(commands)
	} else {
		// part B
	}
}
