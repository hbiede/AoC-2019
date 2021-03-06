package main

import (
    "bufio"
    "flag"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"

    "IntCode"
)

var (
    inputFile = flag.String("inputFile", "inputs/day02.txt", "Input File")
)

const DesiredOutput = 19690720 // cleverly the day of the moon landing

//noinspection GoNilness
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

    commandsClone := make([]int, len(commands))
    copy(commandsClone, commands)

    // part A
    commandsClone[1] = 12
    commandsClone[2] = 2
    fmt.Printf("commands[0] = %d\n", IntCode.IntOpCodeComputer(commandsClone))

    // part B
    for i := 0; i < 100; i++ {
        for j := 0; j < 100; j++ {
            copy(commandsClone, commands)
            commandsClone[1] = i
            commandsClone[2] = j
            if IntCode.IntOpCodeComputer(commandsClone) == DesiredOutput {
                fmt.Printf("Noun: %d\tVerb: %d\n100 * noun + verb (expected 9342): %d", i, j, i*100+j)
            }
        }
    }

}
