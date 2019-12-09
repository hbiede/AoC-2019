package IntCode

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Stream struct {
	Input  chan int
	Output chan int
}

func New() *Stream {
	return &Stream{
		Input:  make(chan int, 1),
		Output: make(chan int),
	}
}

func (stream *Stream) Run(commands []int) {
	IntOpCodeComputerStream(commands, stream)
}

func IntOpCodeComputerStream(commands []int, stream *Stream) {
	// Copy the program into memory, so that we do not modify the original.
	memory := make([]int, len(commands))
	copy(memory, commands)

	getMemoryPointer := func(index int) *int {
		// Grow memory, if index is out of range.
		for len(memory) <= index {
			memory = append(memory, 0)
		}
		return &memory[index]
	}

	var ip, relativeBase int
Processor:
	for {
		instruction := memory[ip]
		opCode := instruction % 100

		getParameter := func(offset int) *int {
			parameter := memory[ip+offset]
			mode := instruction / pow(10, offset+1) % 10
			switch mode {
			case 0: // position mode
				return getMemoryPointer(parameter)
			case 1: // immediate mode
				return &parameter
			case 2: // relative mode
				return getMemoryPointer(relativeBase + parameter)
			default:
				panic(fmt.Sprintf("fault: invalid parameter mode: ip=%d instruction=%d offset=%d mode=%d", ip, instruction, offset, mode))
			}
		}

		switch opCode {
		case 1: // ADD
			a, b, c := getParameter(1), getParameter(2), getParameter(3)
			*c = *a + *b
			ip += 4

		case 2: // MULTIPLY
			a, b, c := getParameter(1), getParameter(2), getParameter(3)
			*c = *a * *b
			ip += 4

		case 3: // INPUT
			*getParameter(1) = <- stream.Input
			ip += 2

		case 4: // OUTPUT
			stream.Output <- *getParameter(1)
			ip += 2

		case 5: // JUMP IF TRUE
			a, b := getParameter(1), getParameter(2)
			if *a != 0 {
				ip = *b
			} else {
				ip += 3
			}

		case 6: // JUMP IF FALSE
			a, b := getParameter(1), getParameter(2)
			if *a == 0 {
				ip = *b
			} else {
				ip += 3
			}

		case 7: // LESS THAN
			a, b, c := getParameter(1), getParameter(2), getParameter(3)
			if *a < *b {
				*c = 1
			} else {
				*c = 0
			}
			ip += 4

		case 8: // EQUAL
			a, b, c := getParameter(1), getParameter(2), getParameter(3)
			if *a == *b {
				*c = 1
			} else {
				*c = 0
			}
			ip += 4

		case 9: // RELATIVE BASE OFFSET
			a := getParameter(1)
			relativeBase += *a
			ip += 2

		case 99: // HALT
			break Processor

		default:
			panic(fmt.Sprintf("fault: invalid opCode: ip=%d instruction=%d opCode=%d", ip, instruction, opCode))
		}
	}
	close(stream.Input)
	close(stream.Output)
}

func pow(a int, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}


/*
 * Legacy code
 */

func IntOpCodeComputer(commands []int) int {
	artificialInput := make([]int, 1)
	returnVal, _ := IntOpCodeComputerProcessor(commands, artificialInput, true)
	return returnVal
}

func IntOpCodeComputerNoPrintWithInput(commands []int, input []int) (int, int) {
	return IntOpCodeComputerProcessor(commands, input, false)
}

/**
 * Returns the value in commands[0] after program completion, -1 if 99 is never reached. Legacy
 */
func IntOpCodeComputerProcessor(commands []int, input []int, print bool) (zeroIndex int, output int) {
	// Process program
	inputIndex := 0
	for i := 0; i < len(commands); i += 0 {
		opCode := commands[i] % 100
		switch opCode {
		case 1:
			commands[commands[i+3]] = getParam(commands, (commands[i]/100)%10, i+1) + getParam(commands, (commands[i]/1000)%10, i+2)
			i += 4
		case 2:
			commands[commands[i+3]] = getParam(commands, (commands[i]/100)%10, i+1) * getParam(commands, (commands[i]/1000)%10, i+2)
			i += 4
		case 3:
			var inputInt int
			if print {
				reader := bufio.NewReader(os.Stdin)
				text, _ := reader.ReadString('\n')
				inputInt, _ = strconv.Atoi(strings.Trim(strings.Replace(text, "\n", "", -1), " "))
			} else {
				inputInt = input[inputIndex%len(input)]
				inputIndex++
			}
			commands[commands[i+1]] = inputInt
			i += 2
		case 4:
			output = getParam(commands, (commands[i]/100)%10, i+1)
			if print {
				fmt.Println(strconv.Itoa(output))
			}
			i += 2
		case 5:
			if getParam(commands, (commands[i]/100)%10, i+1) != 0 {
				i = getParam(commands, (commands[i]/1000)%10, i+2)
			} else {
				i += 3
			}
		case 6:
			if getParam(commands, (commands[i]/100)%10, i+1) == 0 {
				i = getParam(commands, (commands[i]/1000)%10, i+2)
			} else {
				i += 3
			}
		case 7:
			if getParam(commands, (commands[i]/100)%10, i+1) < getParam(commands, (commands[i]/1000)%10, i+2) {
				commands[commands[i+3]] = 1
			} else {
				commands[commands[i+3]] = 0
			}
			i += 4
		case 8:
			if getParam(commands, (commands[i]/100)%10, i+1) == getParam(commands, (commands[i]/1000)%10, i+2) {
				commands[commands[i+3]] = 1
			} else {
				commands[commands[i+3]] = 0
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
