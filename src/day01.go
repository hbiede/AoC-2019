package main

import (
    "bufio"
    "flag"
    "fmt"
    "log"
    "math"
    "os"
    "strconv"
)

var (
    partA     = flag.Bool("partA", false, "Perform part A solution?")
    inputFile = flag.String("inputFile", "inputs/day01.txt", "Input File")
)

func main() {
    flag.Parse()

    file, err := os.Open(*inputFile)
    if err != nil {
        log.Fatal(err)
    }
    //noinspection GoUnhandledErrorResult
    defer file.Close()

    scanner := bufio.NewScanner(file)
    sum := 0
    for scanner.Scan() {
        input, err := strconv.Atoi(scanner.Text())
        if err != nil {
            log.Fatal(err)
        }

        sum += findNecessaryFuel(input, 0)
    }

    fmt.Println("Sum: " + strconv.Itoa(sum))
}

func findNecessaryFuel(x int, acc int) int {
    if *partA {
        // part A
        return int(math.Floor(float64(x/3)) - 2)
    } else {
        // part B
        if x <= 0 {
            return acc
        }

        neededFuel := int(math.Max(math.Floor(float64(x/3))-2, 0))
        return findNecessaryFuel(neededFuel, acc+neededFuel)
    }
}
