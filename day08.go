package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
)

var (
	inputFile   = flag.String("inputFile", "inputs/day08.txt", "Input File")
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

	imageHeight := 6
	imageWidth := 25
	layerComponents := []rune(inputStringFromFile)
	layers := runesToStringArray(layerComponents, imageHeight, imageWidth)

	_, ones, twos := minLayerBreakDown(layers)
	fmt.Printf("Ones times twos in layer with fewest 0's (expected 1716): (%d * %d) = %d", ones, twos, ones * twos)
}

func runesToStringArray(runes []rune, height int, width int) []string {
	imageDim := height * width
	returnArray := make([]string, 0)

	for i := 0; i < len(runes); i += imageDim {
		returnArray = append(returnArray, string(runes[i : i + imageDim])) // slice ranges are [,)
	}
	return returnArray
}

func minLayerBreakDown(layers []string) (zeros int, ones int, twos int) {
	zeros = math.MaxInt64
	for _, layer := range layers {
		zeroCount := 0
		oneCount := 0
		twoCount := 0
		for _, digit := range layer {
			switch digit {
			case '0':
				zeroCount++
			case '1':
				oneCount++
			case '2':
				twoCount++
			}
		}
		if zeroCount < zeros {
			zeros = zeroCount
			ones = oneCount
			twos = twoCount
		}
	}
	return zeros, ones, twos
}