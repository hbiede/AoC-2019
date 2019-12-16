package main

import (
    "bufio"
    "flag"
    "fmt"
    "log"
    "math"
    "os"
    "sort"
)

type Coord struct {
    X int
    Y int
}

func (c1 Coord) angleTo(c2 Coord) float64 {
    angle := math.Atan2(float64(c2.X-c1.X), float64(c1.Y-c2.Y)) * 180 / math.Pi
    if angle < 0 {
        angle += 360
    }
    if angle >= 360 {
        angle -= 360
    }
    return angle
}

func (c1 Coord) distanceTo(c2 Coord) float64 {
    return math.Sqrt(math.Pow(float64(c1.X-c2.X), 2) + math.Pow(float64(c1.Y-c2.Y), 2))
}

func (c1 Coord) equals(c2 Coord) bool {
    return c1.X == c2.X && c1.Y == c2.Y
}

var (
    inputFile = flag.String("inputFile", "inputs/day10.txt", "Input File")
    asteroid  = '#'
    toDestroy = 200
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
    var asteroids []Coord
    for y := 0; scanner.Scan(); y++ {
        for x, char := range scanner.Text() {
            if char == asteroid {
                asteroids = append(asteroids, Coord{X: x, Y: y})
            }
        }
    }

    bestFit, canSee := BestLocation(asteroids)
    fmt.Printf("(%d, %d) can see %d asteroids (Expected (14, 17) can see 260)\n", bestFit.X, bestFit.Y, canSee)

    lastDestroyed := destroyAsteroids(asteroids, bestFit, toDestroy)
    if bestFit.equals(lastDestroyed) {
        log.Fatalln("problem: Got the laser station was destroyed. We're now floating in space")
    } else {
        var ordinalSuffix string
        //noinspection GoBoolExpressions - It may be constant, but this future proofs it
        if toDestroy%10 == 1 && toDestroy%100 != 11 {
            ordinalSuffix = "st"
        } else if toDestroy%10 == 2 {
            ordinalSuffix = "nd"
        } else if toDestroy%10 == 3 {
            ordinalSuffix = "rd"
        } else {
            ordinalSuffix = "th"
        }
        fmt.Printf("%d%s asteroid was destroyed at (%d, %d). Expected at (6, 8)\n", toDestroy, ordinalSuffix, lastDestroyed.X, lastDestroyed.Y)
    }
}

func BestLocation(asteroids []Coord) (Coord, int) {
    bestCount := math.MinInt64
    var best Coord
    for _, current := range asteroids {
        count := 0
        angles := make(map[float64]int)

        for _, viewChecker := range asteroids {
            if current != viewChecker {
                angle := current.angleTo(viewChecker)

                if _, keyExists := angles[angle]; !keyExists {
                    // count up if another asteroid on this view angleTo hasn't been found
                    angles[angle] = 1
                    count++
                }
            }
        }

        if count > bestCount {
            bestCount = count
            best = current
        }
    }

    return best, bestCount
}

func getSlopeList(asteroids []Coord, station Coord) []float64 {
    angles := make(map[float64]int)

    for _, asteroid := range asteroids {
        angle := station.angleTo(asteroid)
        if _, keyExists := angles[angle]; !keyExists {
            angles[angle] = 1
        }
    }
    angleSlice := make([]float64, 0)
    for key := range angles {
        angleSlice = append(angleSlice, key)
    }
    sort.Float64s(angleSlice)
    return angleSlice
}

func asteroidsToAngleSlices(asteroids []Coord, laserStation Coord) map[float64][]Coord {
    angleMap := make(map[float64][]Coord)

    for _, asteroid := range asteroids {
        angle := laserStation.angleTo(asteroid)
        if _, keyExists := angleMap[angle]; !keyExists {
            angleMap[angle] = make([]Coord, 0)
        }
        angleMap[angle] = append(angleMap[angle], asteroid)
    }

    for _, value := range angleMap {
        sort.Slice(value, func(i, j int) bool {
            return laserStation.distanceTo(value[i]) < laserStation.distanceTo(value[j])
        })
    }
    return angleMap
}

func destroyAsteroids(asteroids []Coord, laserStation Coord, asteroidsToDestroy int) Coord {
    if len(asteroids) < asteroidsToDestroy {
        return laserStation
    }

    angleMap := asteroidsToAngleSlices(asteroids, laserStation)
    slopeList := getSlopeList(asteroids, laserStation)

    destroyed := 0
    for _, slope := range slopeList {
        if _, slopeExists := angleMap[slope]; slopeExists && len(angleMap[slope]) > 0 {
            destroyed++
            if asteroidsToDestroy == destroyed {
                // this is the last asteroid to destroy
                return angleMap[slope][0]
            }
            angleMap[slope] = angleMap[slope][1:] // "Destroy" the asteroid
        }
    }
    return laserStation
}
