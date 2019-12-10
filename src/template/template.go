package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	partA       = flag.Bool("partA", true, "Perform part A solution?")
	inputFile   = flag.String("inputFile", "inputs/day!DAY!.txt", "Input File")
	inputString = flag.String("inputs", "", "Input string")
	debug       = flag.Bool("debug", false, "Debug?")
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

	if *partA {
		// part A
		fmt.Printf("Part A")
	} else {
		// part B
	}
}