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

	bounds := strings.Split(inputStringFromFile, "-")
	lowerBound, err := strconv.Atoi(bounds[0])
	if err != nil {
		log.Fatal(err)
	}

	upperBound, err := strconv.Atoi(bounds[1])
	if err != nil {
		log.Fatal(err)
	}

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

	fmt.Printf("Valid passwords for part A: %d\nValid passwords for part B: %d", validPasswordsA, validPasswordsB)
}

func validateNumber(digits []int) (partA bool, partB bool) {
	adjacencyLength := 0
	lastDigit := -1
	for _, digit := range digits {
		if digit < lastDigit {
			// decrease
			return false, false
		} else if digit == lastDigit {
			adjacencyLength++
			partA = true
		} else {
			// increase
			if adjacencyLength == 1 {
				partB = true
			}
			adjacencyLength = 0
		}

		lastDigit = digit
	}

	return partA, partB || adjacencyLength == 1
}

func numberToDigits(input int) []int {
	returnVal := make([]int, 0)
	for input > 0 {
		returnVal = append([]int{input % 10}, returnVal...) // Prepend int to the slice
		input /= 10
	}
	return returnVal
}
