package main

import (
    "IntCode"
    "bufio"
    "flag"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
    "sync"
)

var (
    inputFile = flag.String("inputFile", "inputs/day21.txt", "Input File")
    debug     = flag.Bool("debug", false, "Should I print all the debug info?")
)

func main() {
    flag.Parse()

    commands := processInput()
    run(commands, "NOT A J\nNOT B T\nOR T J\nNOT C T\nOR T J\nAND D J\nWALK\n", 19359969)
	run(commands, "NOT A J\nNOT B T\nOR T J\nNOT C T\nOR T J\nAND D J\nNOT E T\nNOT T T\nOR H T\nAND T J\nRUN\n", 1140082748)
}

func processInput() []int {
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
    for _, commandString := range commandStrings {
        command, err := strconv.Atoi(commandString)
        if err != nil {
            log.Fatal(err)
        }
        commands = append(commands, command)
    }
    return commands
}

func run(commands []int, springProgram string, expected int) {
    botController := IntCode.New()
    commandsClone := make([]int, len(commands))
    copy(commandsClone, commands)
    go botController.Run(commandsClone)

    wg := sync.WaitGroup{}
    wg.Add(1)
    go func(botController *IntCode.Stream) {
        defer func() {
            wg.Done()
            recovery := recover()
            if recovery != nil && fmt.Sprint(recovery) != "send on closed channel" {
                panic(recovery)
            }
        }()

        for output := range botController.Output {
            if output > 127 {
                fmt.Printf("%d hull damage (%d expected)\n", output, expected)
            } else {
	            fmt.Printf("%c", output)
            }
        }
    }(botController)
    passFunctionToBot(botController, springProgram)
    wg.Wait()
}

func passFunctionToBot(botController *IntCode.Stream, function string) {
    for _, character := range function {
        botController.Input <- int(character)
    }
}
