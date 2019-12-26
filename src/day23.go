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

type Packet struct {
    X int
    Y int
}

type Queue struct {
    data []Packet
}

func newQueue() *Queue {
    return &Queue{}
}

func (q *Queue) isEmpty() bool {
    return len((*q).data) == 0
}

func (q *Queue) Pop() Packet {
    defer func() {
        if q.data == nil || len(q.data) <= 1 {
            q.data = make([]Packet, 0)
        } else {
            q.data = q.data[1:]
        }
    }()

    return q.Peek()
}

func (q *Queue) Peek() Packet {
    if q.isEmpty() {
        return Packet{
            X: -1,
            Y: -1,
        }
    }
    return q.data[0]
}

func (q *Queue) Push(p Packet) {
    q.data = append(q.data, p)
}

const (
    COMPUTERS = 50
)

var (
    inputFile = flag.String("inputFile", "inputs/day23.txt", "Input File")
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
    computers := make([]*IntCode.Stream, 0)
    packetQueues := make(map[int]*Queue)
    for i := 0; i < COMPUTERS; i++ {
        computers = append(computers, IntCode.New())
        commandsClone := make([]int, len(commands))
        copy(commandsClone, commands)
        go computers[i].Run(commandsClone)
        computers[i].Input <- i

        packetQueues[i] = newQueue()
    }

    first255 := true
    lastSent := Packet{
        X: 2,
        Y: -1,
    }
    NAT := Packet{
        X: -1,
        Y: -1,
    }
    for {
        idle := 0
        for i := 0; i < COMPUTERS; i++ {
            sendRealPacket := !packetQueues[i].isEmpty()
            sendPacket := packetQueues[i].Peek() // defaults to {-1, -1} if empty

            select {
            case destination := <-computers[i].Output:
            	if destination == 255 {
            		NAT = Packet{X: <-computers[i].Output, Y: <-computers[i].Output}
	            } else {
		            packetQueues[destination].Push(Packet{X: <-computers[i].Output, Y: <-computers[i].Output})
		            //fmt.Printf("%d messaging %d\n", i, destination)
	            }
            case computers[i].Input <- sendPacket.X:
            	if sendRealPacket {
		            computers[i].Input <- packetQueues[i].Pop().Y
	            } else {
	                idle++
                }
            }
        }

        if NAT.Y != -1 {
            if first255 {
                fmt.Printf("First Y: %d (Expected 18513)\n", NAT.Y)
                first255 = false
            }
        }


        if idle == COMPUTERS && !first255 && NAT.Y != -1 {
            fmt.Printf("Sending %d\n", NAT.Y)
            if NAT.Y == lastSent.Y {
                fmt.Printf("Double sent Y: %d (expected XXX)\n", lastSent.Y)
                return
                // 14091 is too high
                // 13278 is too low
            } else {
                lastSent = NAT
                packetQueues[0].Push(NAT)
            }
        }
    }
}
