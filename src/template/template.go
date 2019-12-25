package main

import (
	"bufio"
	"flag"
	"log"
	"os"
)

var (
	inputFile   = flag.String("inputFile", "inputs/day!DAY!.txt", "Input File")
)

func main() {
	flag.Parse()

	inputString := processInput()
}

func processInput() string {
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
	return inputStringFromFile
}