package IntCode

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type IntCodeStream struct {
	Input  chan int
	Output chan int
}

func New() *IntCodeStream {
	return &IntCodeStream{
		Input:  make(chan int, 1),
		Output: make(chan int),
	}
}

func IntOpCodeComputerStream(commands []int, stream *IntCodeStream) {
	// Process program
	Processor: for i := 0; i < len(commands); i += 0 {
		opCode := commands[i] % 100
		switch opCode {
		case 1:
			commands[commands[i + 3]] = getParam(commands, (commands[i] / 100) % 10, i + 1) + getParam(commands, (commands[i] / 1000) % 10, i + 2)
			i += 4
		case 2:
			commands[commands[i + 3]] = getParam(commands, (commands[i] / 100) % 10, i + 1) * getParam(commands, (commands[i] / 1000) % 10, i + 2)
			i += 4
		case 3:
			commands[commands[i + 1]] = <- stream.Input
			i += 2
		case 4:
			output := getParam(commands, (commands[i] / 100) % 10, i + 1)
			stream.Output <- output
			i += 2
		case 5:
			if getParam(commands, (commands[i] / 100) % 10, i + 1) != 0 {
				i = getParam(commands, (commands[i] / 1000) % 10, i + 2)
			} else {
				i += 3
			}
		case 6:
			if getParam(commands, (commands[i] / 100) % 10, i + 1) == 0 {
				i = getParam(commands, (commands[i] / 1000) % 10, i + 2)
			} else {
				i += 3
			}
		case 7:
			if getParam(commands, (commands[i] / 100) % 10, i + 1) < getParam(commands, (commands[i] / 1000) % 10, i + 2) {
				commands[commands[i + 3]] = 1
			} else {
				commands[commands[i + 3]] = 0
			}
			i += 4
		case 8:
			if getParam(commands, (commands[i] / 100) % 10, i + 1) == getParam(commands, (commands[i] / 1000) % 10, i + 2) {
				commands[commands[i + 3]] = 1
			} else {
				commands[commands[i + 3]] = 0
			}
			i += 4
		case 99:
			break Processor
		default:
			log.Fatalf("%d is an unknown command\n", commands[i])
		}
	}
	close(stream.Input)
	close(stream.Output)
}

func IntOpCodeComputer(commands []int) int {
	artificialInput := make ([]int, 1)
	returnVal, _ := IntOpCodeComputerProcessor(commands, artificialInput, true)
	return returnVal
}

func IntOpCodeComputerNoPrintWithInput(commands []int, input []int) (int, int) {
	return IntOpCodeComputerProcessor(commands, input, false)
}

/**
 * Returns the value in commands[0] after program completion, -1 if 99 is never reached
 */
func IntOpCodeComputerProcessor(commands []int, input []int, print bool) (zeroIndex int, output int) {
	// Process program
	inputIndex := 0
	for i := 0; i < len(commands); i += 0 {
		opCode := commands[i] % 100
		switch opCode {
		case 1:
			commands[commands[i + 3]] = getParam(commands, (commands[i] / 100) % 10, i + 1) + getParam(commands, (commands[i] / 1000) % 10, i + 2)
			i += 4
		case 2:
			commands[commands[i + 3]] = getParam(commands, (commands[i] / 100) % 10, i + 1) * getParam(commands, (commands[i] / 1000) % 10, i + 2)
			i += 4
		case 3:
			var inputInt int
			if print {
				reader := bufio.NewReader(os.Stdin)
				text, _ := reader.ReadString('\n')
				inputInt, _ = strconv.Atoi(strings.Trim(strings.Replace(text, "\n", "", -1), " "))
			} else {
				inputInt = input[inputIndex % len(input)]
				inputIndex++
			}
			commands[commands[i + 1]] = inputInt
			i += 2
		case 4:
			output = getParam(commands, (commands[i] / 100) % 10, i + 1)
			if print {
				fmt.Println(strconv.Itoa(output))
			}
			i += 2
		case 5:
			if getParam(commands, (commands[i] / 100) % 10, i + 1) != 0 {
				i = getParam(commands, (commands[i] / 1000) % 10, i + 2)
			} else {
				i += 3
			}
		case 6:
			if getParam(commands, (commands[i] / 100) % 10, i + 1) == 0 {
				i = getParam(commands, (commands[i] / 1000) % 10, i + 2)
			} else {
				i += 3
			}
		case 7:
			if getParam(commands, (commands[i] / 100) % 10, i + 1) < getParam(commands, (commands[i] / 1000) % 10, i + 2) {
				commands[commands[i + 3]] = 1
			} else {
				commands[commands[i + 3]] = 0
			}
			i += 4
		case 8:
			if getParam(commands, (commands[i] / 100) % 10, i + 1) == getParam(commands, (commands[i] / 1000) % 10, i + 2) {
				commands[commands[i + 3]] = 1
			} else {
				commands[commands[i + 3]] = 0
			}
			i += 4
		case 99:
			return commands[0], output
		default:
			log.Fatalf("%d is an unknown command\n", commands[i])
		}
	}
	return -1, -1
}

func getParam(commands []int, immediacy int, index int) int {
	if immediacy == 1 {
		return commands[index]
	} else {
		return commands[commands[index]]
	}
}
