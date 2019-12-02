package main

import (
	"log"
)

/**
 * Returns the value in commands[0] after program completion, -1 if 99 is never reached
 */
func IntOpCodeComputer(commands []int) int {
	// Process program
	for i := 0; i + 3 < len(commands); i += 4 {
		switch commands[i] {
		case 1:
			commands[commands[i + 3]] = commands[commands[i + 1]] + commands[commands[i + 2]]
		case 2:
			commands[commands[i + 3]] = commands[commands[i + 1]] * commands[commands[i + 2]]
		case 99:
			return commands[0]
		default:
			log.Fatalf("%d is an unknown command\n", commands[i])
		}
	}
	return -1
}
