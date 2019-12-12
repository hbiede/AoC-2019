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
	inputFile = flag.String("inputFile", "inputs/day04.txt", "Input File")
)

func main() {
	flag.Parse()

	lowerBound, upperBound := processInput()

	validPasswordsA := 0
	validPasswordsB := 0
	for i := lowerBound; i <= upperBound; i++ {
		partAValid, partBValid := validateNumber(numberToDigits(i))
		if partAValid {
			validPasswordsA++
		}
		if partBValid {
			validPasswordsB++
		}
	}

	fmt.Printf("Valid passwords for part A (expected 1246): %d\nValid passwords for part B (expected 814): %d", validPasswordsA, validPasswordsB)
}

func processInput() (int, int) {
	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	//noinspection GoUnhandledErrorResult
	defer file.Close()

	// Read in inputs
	scanner := bufio.NewScanner(file)
	inputStringFromFile := ""
	scanner.Scan()
	inputStringFromFile += scanner.Text()

	bounds := strings.Split(inputStringFromFile, "-")
	lowerBound, err := strconv.Atoi(bounds[0])
	if err != nil {
		log.Fatal(err)
	}

	upperBound, err := strconv.Atoi(bounds[1])
	if err != nil {
		log.Fatal(err)
	}

	return lowerBound, upperBound
}

func validateNumber(digits []int) (partA bool, partB bool) {
	adjacencyLength := 1
	lastDigit := -1
	for _, digit := range digits {
		if digit < lastDigit {
			// decreasing values
			return false, false
		} else if digit == lastDigit {
			// repeating values
			adjacencyLength++
			partA = true
		} else {
			// increasing values: ✅
			if adjacencyLength == 2 {
				// only say the repeating rule in partB is true if it repeats once
				partB = true
			}
			adjacencyLength = 1
		}

		lastDigit = digit
	}

	return partA, partB || adjacencyLength == 2
}

func numberToDigits(input int) []int {
	returnVal := make([]int, 0)
	for input > 0 {
		returnVal = append([]int{input % 10}, returnVal...) // Prepend int to the slice
		input /= 10
	}
	return returnVal
}
