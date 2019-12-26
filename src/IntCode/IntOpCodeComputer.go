package IntCode

import (
    "fmt"
)

type Stream struct {
    Input     chan int
    Output    chan int
    IsRunning bool
}

func New() *Stream {
    return &Stream{
        Input:  make(chan int, 100),
        Output: make(chan int, 100),
    }
}

var ipJumps = map[int]int{
    1: 4,
    2: 4,
    3: 2,
    4: 2,
    5: 3,
    6: 3,
    7: 4,
    8: 4,
    9: 2,
}

func (stream *Stream) Run(commands []int) {
	stream.IsRunning = true
    IntOpCodeComputerStream(commands, stream)
}

func IntOpCodeComputerStream(commands []int, stream *Stream) {
    // Copy the program into memory, so that we do not modify the original.
    memory := make([]int, len(commands))
    copy(memory, commands)

    getMemoryPointer := func(index int) *int {
        // Grow memory, if index is out of range.
        for len(memory) <= index {
            memory = append(memory, 0)
        }
        return &memory[index]
    }

    relativeBase := 0
Processor:
    for ip := 0; ip < len(memory); ip += 0 {
        instruction := memory[ip]
        opCode := instruction % 100

        getParameter := func(offset int) *int {
            parameter := memory[ip+offset]
            mode := instruction / pow(10, offset+1) % 10
            switch mode {
            case 0: // position mode
                return getMemoryPointer(parameter)
            case 1: // immediate mode
                return &parameter
            case 2: // relative mode
                return getMemoryPointer(relativeBase + parameter)
            default:
                panic(fmt.Sprintf("fault: invalid parameter mode: ip=%d instruction=%d offset=%d mode=%d", ip, instruction, offset, mode))
            }
        }

        switch opCode {
        case 1: // ADD
            a, b, c := getParameter(1), getParameter(2), getParameter(3)
            *c = *a + *b
        case 2: // MULTIPLY
            a, b, c := getParameter(1), getParameter(2), getParameter(3)
            *c = *a * *b
        case 3: // INPUT
            *getParameter(1) = <-stream.Input
        case 4: // OUTPUT
            stream.Output <- *getParameter(1)
        case 5: // JUMP IF TRUE
            a, b := getParameter(1), getParameter(2)
            if *a != 0 {
                ip = *b
                continue
            }
        case 6: // JUMP IF FALSE
            a, b := getParameter(1), getParameter(2)
            if *a == 0 {
                ip = *b
                continue
            }
        case 7: // LESS THAN
            a, b, c := getParameter(1), getParameter(2), getParameter(3)
            if *a < *b {
                *c = 1
            } else {
                *c = 0
            }
        case 8: // EQUAL
            a, b, c := getParameter(1), getParameter(2), getParameter(3)
            if *a == *b {
                *c = 1
            } else {
                *c = 0
            }
        case 9: // RELATIVE BASE OFFSET
            a := getParameter(1)
            relativeBase += *a
        case 99: // HALT
            stream.IsRunning = false
            break Processor
        default:
            panic(fmt.Sprintf("fault: invalid opCode: ip=%d instruction=%d opCode=%d", ip, instruction, opCode))
        }
        ip += ipJumps[opCode]
    }
    close(stream.Input)
    close(stream.Output)
}

func pow(a int, b int) int {
    p := 1
    for b > 0 {
        if b&1 != 0 {
            p *= a
        }
        b >>= 1
        a *= a
    }
    return p
}
