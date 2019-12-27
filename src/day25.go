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
    inputFile = flag.String("inputFile", "inputs/day25.txt", "Input File")
)

func main() {
    flag.Parse()

    commands := processInput()
    run(commands)
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
    for _, commandString := range commandStrings { // the _ disregards the index and keeps the element in commandString
        command, err := strconv.Atoi(commandString)
        if err != nil {
            log.Fatal(err)
        }
        commands = append(commands, command)
    }
    return commands
}

func run(commands []int) {
    gameController := IntCode.New()
    commandsClone := make([]int, len(commands))
    copy(commandsClone, commands)
    go gameController.Run(commandsClone)

    wg := sync.WaitGroup{}

    wg.Add(1)
    go func() {
        defer func() {
            wg.Done()
            recovery := recover()
            if recovery != nil && fmt.Sprint(recovery) != "send on closed channel" {
                panic(recovery)
            }
        }()

        for output := range gameController.Output {
            fmt.Printf("%c", output)
        }
    }()

    go func() {
    	for {
		    reader := bufio.NewReader(os.Stdin)
		    text, _ := reader.ReadString('\n')
		    passFunctionToBot(gameController, text)
	    }
    }()
    wg.Wait()
}

func passFunctionToBot(botController *IntCode.Stream, function string) {
    for _, character := range function {
        botController.Input <- int(character)
    }
}
