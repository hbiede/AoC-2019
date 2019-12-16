package main

import (
    "bufio"
    "flag"
    "fmt"
    "image"
    "image/color"
    "image/draw"
    "image/png"
    "log"
    "math"
    "os"
)

var (
    inputFile = flag.String("inputFile", "inputs/day08.txt", "Input File")
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

    imageHeight := 6
    imageWidth := 25
    layerComponents := []rune(inputStringFromFile)
    layers := runesToStringArray(layerComponents, imageHeight, imageWidth)

    _, ones, twos := minLayerBreakDown(layers)
    fmt.Printf("Ones times twos in layer with fewest 0's (expected 1716): (%d * %d) = %d\n", ones, twos, ones*twos)

    f, _ := os.Create("outputs/day08-2.png")
    _ = png.Encode(f, createImage(reverse(layers), imageHeight, imageWidth))
    fmt.Println("Image generated. KFABY expected")

}

func runesToStringArray(runes []rune, height int, width int) []string {
    imageDim := height * width
    returnArray := make([]string, 0)

    for i := 0; i < len(runes); i += imageDim {
        returnArray = append(returnArray, string(runes[i:i+imageDim])) // slice ranges are [,)
    }
    return returnArray
}

func minLayerBreakDown(layers []string) (zeros int, ones int, twos int) {
    zeros = math.MaxInt64
    for _, layer := range layers {
        zeroCount := 0
        oneCount := 0
        twoCount := 0
        for _, digit := range layer {
            switch digit {
            case '0':
                zeroCount++
            case '1':
                oneCount++
            case '2':
                twoCount++
            }
        }
        if zeroCount < zeros {
            zeros = zeroCount
            ones = oneCount
            twos = twoCount
        }
    }
    return zeros, ones, twos
}

func createImage(layers []string, height int, width int) draw.Image {
    topLeft := image.Point{X: 0, Y: 0}
    bottomRight := image.Point{X: width, Y: height}
    img := image.NewRGBA(image.Rectangle{Min: topLeft, Max: bottomRight})
    white := color.White
    black := color.Black
    for _, layer := range layers {
        for i, character := range layer {
            var setColor color.Color
            switch character {
            case '0':
                setColor = black
            case '1':
                setColor = white
            default:
                setColor = nil
            }
            if setColor != nil {
                img.Set(i%width, i/width, setColor)
            }
        }
    }
    return img
}

func reverse(array []string) []string {
    arrayCopy := make([]string, len(array))
    for left, right := 0, len(array)-1; left < right; left, right = left+1, right-1 {
        arrayCopy[left], arrayCopy[right] = array[right], array[left]
    }
    return arrayCopy
}
