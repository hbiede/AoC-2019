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
)

var (
    inputFile   = flag.String("inputFile", "inputs/day16.txt", "Input File")
    phases      = flag.Int("phases", 100, "Number of phases to run in the given pattern")
    repeatCount = flag.Int("repeatCount", 10000, "Number of times the signal repeats for part 2")
)

func main() {
    flag.Parse()

    inputStringFromFile := processInput()
    valuesA := decodeSignal(inputStringFromFile, fftPhase1)
    for i := 0; i < 8 && i < len(valuesA); i++ {
        fmt.Printf("%d", valuesA[i])
    }
    fmt.Println(" (Expected 74369033)")

    valuesB := decodeSignal(strings.Repeat(inputStringFromFile, *repeatCount), fftPhase2)
    messageLocation, err := strconv.Atoi(inputStringFromFile[0:7])
    //valuesB := decodeSignal(strings.Repeat("03036732577212944063491565474664", *repeatCount))
    //messageLocation, err := strconv.Atoi("03036732577212944063491565474664"[0:7])
    if err != nil {
        log.Fatalf("%s is not an integer\n", inputStringFromFile[0:7])
    }
    for _, digit := range valuesB[messageLocation: messageLocation + 8] {
        fmt.Printf("%d", digit)
    }
    fmt.Println(" (Expected 19903864)")
}

func decodeSignal(signal string, whichPart func([]int) []int) []int {
    values := stringToDigits(signal)
    for i := 0; i < *phases; i++ {
        values = whichPart(values)
    }
    return values
}

func processInput() string {
    file, err := os.Open(*inputFile)
    if err != nil {
        log.Fatal(err)
    }
    //noinspection GoUnhandledErrorResult
    defer file.Close()

    // Read in inputs
    scanner := bufio.NewScanner(file)
    scanner.Scan()

    return scanner.Text()
}

func stringToDigits(input string) []int {
    returnVal := make([]int, 0)
    for _, r := range input {
        // dirty way of converting strings to digits. I trust the user to give valid input
        returnVal = append(returnVal, int(r - '0'))
    }
    return returnVal
}

func fftPhase1(values []int) []int {
    pattern := []int{0, 1, 0, -1}
    returnVal := make([]int, len(values))
    for i := range values {
        for j, value := range values {
            patternVal := pattern[((j + 1) / (i + 1)) % len(pattern)]
            returnVal[i] += value * patternVal
        }
        returnVal[i] = int(math.Abs(float64(returnVal[i] % 10)))
    }
    return returnVal
}

func fftPhase2(values []int) []int {
    // Initial thoughts:
    // to be honest, I'm not 100% certain how this works. I was trying to piece something together when my first
    // solution was taking forever to complete one phase, and this idea worked for the back half of the part 1 examples,
    // so I hoped it would work for enough of this large phase to get the message out, and it did.
    // It doesn't work for part 1 unfortunately (it would be much faster to use this O(n) alg over than O(n^2) one).
    // I'm guessing it only works for this one because we aren't using the front 8 digits and this relies on jumping to
    // the middle via the 'messageLocation' idea?

    // After further thought:
    // This works because, in the second half of the array, the only part of the pattern being used is the 0 and the 1.
    // Since the 0 will eliminate all the values before a given index, the last index will stay constant, and then every
    // index before that would just be the sum of all values after it. And since only the last digit matters in that
    // addition, you can just use the index after the index you're looking at (since it already encompasses the sum of
    // everything after it). Since we assume the messageLocation > the midpoint, we can safely ignore anything before it
    messageLocation := 0
    const messageLength = 7
    for i := 0; i < messageLength; i++ {
        messageLocation += values[messageLength - i - 1] * int(math.Pow(10, float64(i)))
    }

    returnVal := make([]int, len(values))
    returnVal[len(values) - 1] = values[len(values) - 1]

    for i := len(values) - 2; i > messageLocation; i-- {
        returnVal[i] = (values[i] + returnVal[i + 1]) % 10
    }
    return returnVal
}
