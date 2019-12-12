package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type moon struct {
	x    int
	y    int
	z    int
	xVel int
	yVel int
	zVel int
}

func (m *moon) xEquals(other *moon) bool {
	return m.x == other.x && m.xVel == other.xVel
}

func (m *moon) yEquals(other *moon) bool {
	return m.y == other.y && m.yVel == other.yVel
}

func (m *moon) zEquals(other *moon) bool {
	return m.z == other.z && m.zVel == other.zVel
}

func (m *moon) toString() string {
	return fmt.Sprintf("pos=<x=%3d, y=%3d, z=%3d>, vel=<x=%3d, y=%3d, z=%3d>", m.x, m.y, m.z,
		m.xVel, m.yVel, m.zVel)
}

func (m *moon) clone() *moon {
	return &moon{x: m.x, y: m.y, z: m.z, xVel: m.xVel, yVel: m.yVel, zVel: m.zVel}
}

var (
	inputFile = flag.String("inputFile", "inputs/day12.txt", "Input File")
	debug = flag.Bool("debug", false, "Debug Mode")
)

func main() {
	flag.Parse()
	moons := generateMoons(processInput())
	moonOrigins := make([]*moon, len(moons))
	deepCopy(moonOrigins, moons)

	simulate(1000, moons, moonOrigins)

	totalEnergy := 0
	for _, moon := range moons {
		totalEnergy += potentialEnergy(*moon) * kineticEnergy(*moon)
	}

	fmt.Printf("Total Energy in the system (expected 6227): %d\n", totalEnergy)

	deepCopy(moons, moonOrigins) // get the original map back
	simulate(-1, moons, moonOrigins)
}

func deepCopy(dst []*moon, org []*moon) {
	for i, moon := range org {
		dst[i] = moon.clone()
	}
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
		inputStringFromFile += scanner.Text() + "\n"
	}
	inputStringFromFile = strings.TrimSuffix(inputStringFromFile, "\n")

	inputStringFromFile = strings.ReplaceAll(inputStringFromFile, "<x=", "")
	inputStringFromFile = strings.ReplaceAll(inputStringFromFile, " y=", "")
	inputStringFromFile = strings.ReplaceAll(inputStringFromFile, " z=", "")
	inputStringFromFile = strings.ReplaceAll(inputStringFromFile, ">", "")
	return inputStringFromFile
}

func generateMoons(inputStringFromFile string) []*moon {
	planetStrings := strings.Split(inputStringFromFile, "\n")
	moons := make([]*moon, 0)
	for _, planetString := range planetStrings {
		planetLoc := strings.Split(planetString, ",")
		xLoc, errX := strconv.Atoi(planetLoc[0])
		yLoc, errY := strconv.Atoi(planetLoc[1])
		zLoc, errZ := strconv.Atoi(planetLoc[2])
		if errX != nil {
			log.Fatalln(errX)
		}

		if errY != nil {
			log.Fatalln(errY)
		}

		if errZ != nil {
			log.Fatalln(errZ)
		}
		moons = append(moons, &moon{x: xLoc, y: yLoc, z: zLoc})
	}
	return moons
}

func simulate(steps int, moons []*moon, moonOrigins []*moon) {
	i := 0
	xRepeatRate, yRepeatRate, zRepeatRate := -1, -1, -1
	for {
		if i == steps {
			break
		}
		if *debug {
			fmt.Println(strconv.Itoa(i+1))
		}
		stepAllVelocities(moons)
		movePlanets(moons)
		i++

		// 346865676208065 - too high
		// 325192689685472 - too low
		xRepeated, yRepeated, zRepeated := dimensionsAreEquivalent(moons, moonOrigins)
		if xRepeated && xRepeatRate == -1 {
			xRepeatRate = i
		}
		if yRepeated && yRepeatRate == -1 {
			yRepeatRate = i
		}
		if zRepeated && zRepeatRate == -1 {
			zRepeatRate = i
		}
		if xRepeatRate != -1 && yRepeatRate != -1 && zRepeatRate != -1 {
			fmt.Printf("History repeated after %d steps\n", lcm(xRepeatRate, yRepeatRate, zRepeatRate))
			return
		}
	}
}

// greatest common divisor (GCD) via Euclidean algorithm from GoPlayground
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD from GoPlayground
func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

func dimensionsAreEquivalent(current []*moon, original []*moon) (bool, bool, bool) {
	if len(current) != len(original) {
		return false, false, false
	}
	xEquals, yEquals, zEquals := true, true, true
	for i, currentMoon := range current {
		xEquals = currentMoon.xEquals(original[i]) && xEquals
		yEquals = currentMoon.yEquals(original[i]) && yEquals
		zEquals = currentMoon.zEquals(original[i]) && zEquals
	}
	return xEquals, yEquals, zEquals
}

func stepAllVelocities(moons []*moon) {
	for j, moonA := range moons {
		for k, moonB := range moons {
			if j < k {
				updateVelocities(moonA, moonB)
			}
		}
	}
}

func updateVelocities(moonA *moon, moonB *moon) {
	moonA.xVel += signum(moonB.x, moonA.x)
	moonB.xVel += signum(moonA.x, moonB.x)

	moonA.yVel += signum(moonB.y, moonA.y)
	moonB.yVel += signum(moonA.y, moonB.y)

	moonA.zVel += signum(moonB.z, moonA.z)
	moonB.zVel += signum(moonA.z, moonB.z)
}

func signum(x int, x2 int) int {
	if x == x2 {
		return 0
	}
	test := int(math.Abs(float64(x-x2)) / float64(x-x2))
	return test
}

func movePlanets(moons []*moon) {
	for _, moon := range moons {
		moon.x += moon.xVel
		moon.y += moon.yVel
		moon.z += moon.zVel
		if *debug {
			fmt.Println(moon.toString())
		}
	}
	if *debug {
		fmt.Println()
	}
}

func potentialEnergy(moon moon) int {
	return int(math.Abs(float64(moon.x)) + math.Abs(float64(moon.y)) + math.Abs(float64(moon.z)))
}

func kineticEnergy(moon moon) int {
	return int(math.Abs(float64(moon.xVel)) + math.Abs(float64(moon.yVel)) + math.Abs(float64(moon.zVel)))
}

