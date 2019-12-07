package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	inputFile   = flag.String("inputFile", "inputs/day06-official.txt", "Input File")
)

// memoization table
var orbitCount = make(map[string]int)

// maps the string name of the planet onto the name of the entity it orbits
var orbitMap = make(map[string]string)

func main() {
	flag.Parse()

	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	//noinspection GoUnhandledErrorResult
	defer file.Close()

	// Read in inputs and put them in the map
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputString := strings.Split(scanner.Text(),")")
		orbitMap[inputString[1]] = inputString[0]
	}
	// acts as the base case for the memoization table
	orbitCount["COM"] = 0

	// sum all direct and indirect orbits
	orbitSum := 0
	for orbiterPlanet := range orbitMap {
		orbitSum += countOrbits(orbiterPlanet)
	}
	fmt.Printf("Part A: %d orbits\n", orbitSum)

	// thinking of the orbits are as a tree with infinite roots ber node, this finds the nearest common parent node
	planetPathToYou := ""
	currentPlanet := orbitMap["YOU"]
	for currentPlanet != "COM" {
		currentPlanet = orbitMap[currentPlanet]
		planetPathToYou = planetPathToYou + "-" + currentPlanet
	}
	var commonParentOrbit string
	currentPlanet = orbitMap["SAN"]
	for currentPlanet != "COM" {
		if strings.Contains(planetPathToYou, currentPlanet) {
			commonParentOrbit = currentPlanet
			fmt.Println(commonParentOrbit)
			break
		}
		currentPlanet = orbitMap[currentPlanet]
	}

	// Distance from COM to SAN + COM to YOU gives you the distance if you went all the way to the common orbit entity,
	// so, if you subtract twice the distance from COM to their common orbit entity, you get the direct distance between
	// them. Then you have to subtract 2 since we are going between the planets you and santa are orbiting, not you and
	// santa yourselves
	fmt.Printf("Part A: %d transfers\n", orbitCount["YOU"] + orbitCount["SAN"] - 2 * orbitCount[commonParentOrbit] - 2)
}

/**
 * Find the number of orbits for a planet in a memoized fashion
 */
func countOrbits(planet string) int {
	if _, keyExists := orbitCount[planet]; !keyExists {
		// if the orbit count hasn't been found for this planet, find it and add it to the table
		orbitCount[planet] = 1 + countOrbits(orbitMap[planet])
	}
	return orbitCount[planet]
}
