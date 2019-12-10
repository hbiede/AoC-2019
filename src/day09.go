package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	inputFile = flag.String("inputFile", "inputs/day09.txt", "Input File")
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

	// Part A
	boost := IntCode.New()
	commandsClone := make([]int, len(commands))
	copy(commandsClone, commands)
	go boost.Run(commandsClone)

	boost.Input <- 1
	fmt.Println("Expected: 2427443564")
	for output := range boost.Output {
		fmt.Println(strconv.Itoa(output))
	}

	// Part B
	boost = IntCode.New()
	copy(commandsClone, commands)
	go boost.Run(commandsClone)

	boost.Input <- 2
	fmt.Println("Expected: 87221")
	for output := range boost.Output {
		fmt.Println(strconv.Itoa(output))
	}
}
