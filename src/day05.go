package main

import (
    "bufio"
    "flag"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"

    "./IntCode"
)

var (
    inputFile = flag.String("inputFile", "inputs/day05.txt", "Input File")
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
        commands = append(commands, int(command))
    }

    commandsClone := make([]int, len(commands))
    copy(commandsClone, commands)
    artificialInput := make([]int, 1)
    artificialInput[0] = 1
    _, partA := IntCode.IntOpCodeComputerNoPrintWithInput(commandsClone, artificialInput)
    fmt.Printf("Got %d... expected 6731945\n", partA)

    copy(commandsClone, commands)
    artificialInput[0] = 5
    _, partB := IntCode.IntOpCodeComputerNoPrintWithInput(commandsClone, artificialInput)
    fmt.Printf("Got %d... expected 9571668", partB)
}
