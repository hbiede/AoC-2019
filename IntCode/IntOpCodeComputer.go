package IntCode

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/**
 * Returns the value in commands[0] after program completion, -1 if 99 is never reached
 */
func IntOpCodeComputer(commands []int) int {
	// Process program
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
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			input, _ := strconv.Atoi(strings.Trim(strings.Replace(text, "\n", "", -1), " "))
			commands[commands[i + 1]] = input
			i += 2
		case 4:
			fmt.Println(strconv.Itoa(getParam(commands, (commands[i] / 100) % 10, i + 1)))
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
			return commands[0]
		default:
			log.Fatalf("%d is an unknown command\n", commands[i])
		}
	}
	return -1
}

func getParam(commands []int, immediacy int, index int) int {
	if immediacy == 1 {
		return commands[index]
	} else {
		return commands[commands[index]]
	}
}
