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

		sum += findNecessaryFuel(input, 0)
	}

	fmt.Println("Sum: " + strconv.Itoa(sum))
}

func findNecessaryFuel(x int, acc int) int {
	if x <= 0 {
		return acc
	}

	neededFuel := int(math.Max(math.Floor(float64(x/3)) - 2, 0))

	return findNecessaryFuel(neededFuel, acc + neededFuel)
}