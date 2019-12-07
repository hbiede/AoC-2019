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
	"sync"

	"./IntCode"
)

var (
	inputFile   = flag.String("inputFile", "inputs/day07.txt", "Input File")
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

	maxThrust := math.MinInt64
	for i := 0; i <= 4; i++ {
		for j := 0; j <= 4; j++ {
			for k := 0; k <= 4; k++ {
				for l := 0; l <= 4; l++ {
					for m := 0; m <= 4; m++ {
						if i != j && i != k && i != l && i != m && j != k && j != l && j != m && k != l && k != m && l != m {
							ampA := calcThrustAmp(i, 0, commands)
							ampB := calcThrustAmp(j, ampA, commands)
							ampC := calcThrustAmp(k, ampB, commands)
							ampD := calcThrustAmp(l, ampC, commands)
							ampE := calcThrustAmp(m, ampD, commands)
							if ampE > maxThrust {
								maxThrust = ampE
							}
						}
					}
				}
			}
		}
	}

	fmt.Printf("Part 1: Got %d, expected 225056\n", maxThrust)

	maxThrust = math.MinInt64
	for i := 5; i <= 9; i++ {
		for j := 5; j <= 9; j++ {
			for k := 5; k <= 9; k++ {
				for l := 5; l <= 9; l++ {
					for m := 5; m <= 9; m++ {
						if i != j && i != k && i != l && i != m && j != k && j != l && j != m && k != l && k != m && l != m {
							phases := []int{i, j, k, l, m}
							amps := [5]*IntCode.IntCodeStream{
								IntCode.New(),
								IntCode.New(),
								IntCode.New(),
								IntCode.New(),
								IntCode.New(),
							}

							for i, ics := range amps {
								ics.Input <- phases[i]
								commandsClone := make([]int, len(commands))
								copy(commandsClone, commands)
								go IntCode.IntOpCodeComputerStream(commandsClone, ics)
							}
							amps[0].Input <- 0

							output := -1
							wg := sync.WaitGroup{}
							for ampIndex, ics := range amps {
								wg.Add(1)
								go func(ampIndex int, ics *IntCode.IntCodeStream) {
									defer func() {
										wg.Done()
										recovery := recover()
										if recovery != nil && fmt.Sprint(recovery) != "send on closed channel"{
											panic(recovery)
										}
									}()

									for outputThrust := range ics.Output {
										if ampIndex == len(amps) - 1 {
											output = outputThrust
										}
										amps[(ampIndex + 1) % len(amps)].Input <- outputThrust
									}
								}(ampIndex, ics)
							}
							wg.Wait()
							if output > maxThrust {
								maxThrust = output
							}
						}
					}
				}
			}
		}
	}

	fmt.Printf("Part 2: Got %d, expected 14260332\n", maxThrust)
}

func calcThrustAmp(inputA int, inputB int, commands []int) int {
	commandsClone := make([]int, len(commands))
	copy(commandsClone, commands)
	artificialInput := make([]int, 2)
	artificialInput[0] = inputA
	artificialInput[1] = inputB
	_, ampThrust := IntCode.IntOpCodeComputerNoPrintWithInput(commandsClone, artificialInput)
	return ampThrust
}
