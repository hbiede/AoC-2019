package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		input, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		sum += findNecessaryFuel(input)
	}

	fmt.Println("Sum: " + strconv.Itoa(sum))
}

func findNecessaryFuel(x int) int {
	neededFuel := int(math.Max(math.Floor(float64(x/3)) - 2, 0))

	if neededFuel > 0 {
		neededFuel += findNecessaryFuel(neededFuel)
	}
	return neededFuel
}