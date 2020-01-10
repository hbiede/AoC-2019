package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

var (
	inputFile   = flag.String("inputFile", "inputs/day22.txt", "Input File")
	deckLengthA = flag.Int("deckLengthA", 10007, "The size of the deck to be used in part A")
	deckLengthB = flag.Int64("deckLengthB", 119315717514047, "The size of the deck to be used in part B")
	iterationsB = flag.Int64("iterationsB", 101741582076661, "The iterations of shuffling to be used in part B")
)

type shuffleType int

//noinspection GoSnakeCaseUsage - Dumb convention
const (
	DEAL shuffleType = iota + 1 // auto increment constant values from 0 to length - 1
	CUT
	DEAL_N
)

type shuffle struct {
	sType shuffleType
	value int
}


func main() {
	flag.Parse()

	shuffles := processInput()

	deck := make([]int, *deckLengthA)
	for i := range deck {
		deck[i] = i
	}

	shuffleDeck(&deck, shuffles)
	for i, card := range deck {
		if card == 2019 {
			fmt.Printf("2019 is in position %d (1867 expected)\n", i)
			break
		}
	}

	fmt.Printf("%d was in the 2020th position\n", modShuffleDeck(2020, *deckLengthB, *iterationsB, shuffles))
}

func processInput() []shuffle {
	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	//noinspection GoUnhandledErrorResult
	defer file.Close()

	// Read in inputs
	scanner := bufio.NewScanner(file)
	shufflesFromFile := make([]shuffle, 0)
	for scanner.Scan() {
		shufflesFromFile = append(shufflesFromFile, parseShuffleType(scanner.Text()))
	}
	return shufflesFromFile
}

func parseShuffleType(input string) shuffle {
	if strings.Contains(input, "cut") {
		cutLength, err := strconv.Atoi(strings.TrimPrefix(input, "cut "))
		if err != nil {
			log.Fatalf("%s is an invalid line", input)
		}
		return shuffle {
			sType: CUT,
			value: cutLength,
		}
	} else if strings.Contains(input, "increment") {
		increment, err := strconv.Atoi(strings.TrimPrefix(input, "deal with increment "))
		if err != nil {
			log.Fatalf("%s is an invalid line", input)
		}
		return shuffle {
			sType: DEAL_N,
			value: increment,
		}
	} else {
		return shuffle{
			sType: DEAL,
		}
	}
}

func shuffleDeck(deck *[]int, shuffles []shuffle) {
	for _, s := range shuffles {
		switch s.sType {
		case DEAL:
			dealDeck(deck)
		case CUT:
			cutDeck(deck, s.value)
		case DEAL_N:
			dealDeckWithIncrement(deck, s.value)
		}
	}
}

func cutDeck(deck *[]int, size int) {
	if deck == nil || size % len(*deck) == 0 {
		log.Print("Invalid input to cutDeck\n")
		return
	}

	usableSize := size
	for usableSize < 0 {
		usableSize += len(*deck)
	}
	newOrder := append((*deck)[usableSize:], (*deck)[:usableSize]...)

	for i, card := range newOrder {
		(*deck)[i] = card
	}
}

func dealDeck(deck *[]int) {
	reverseDeck(deck)
}

func dealDeckWithIncrement(deck *[]int, increment int) {
	if deck == nil || increment % len(*deck) == 0 || increment % len(*deck) == 1 {
		log.Print("Invalid input to dealDeckWithIncrement\n")
		return
	} else if increment % len(*deck) == -1 {
		reverseDeck(deck)
	} else {
		newOrder := make([]int, len(*deck))

		for i, card := range *deck {
			newOrder[(i * increment) % len(*deck)] = card
		}

		for i, card := range newOrder {
			(*deck)[i] = card
		}
	}
}

func reverseDeck(deck *[]int) {
	for i := len(*deck)/2-1; i >= 0; i-- {
		opposite := len(*deck)-i-1
		(*deck)[i], (*deck)[opposite] = (*deck)[opposite], (*deck)[i]
	}
}

func modShuffleDeck(finalPositionToFind int64, length int64, iterations int64, shuffles []shuffle) int64 {
	// Never would have figured this out solo... but here is my understanding of this concept now:
	// Since you can represent each of the 3 operations as a relatively straight-forward iterative, mathematical
	// functions, and mathematical functions are linear combinations of more basic functions, you can create one large
	// (in this case `finalFormula`), you can take the multipliers and additives (m and b from mx+b), and combine them
	// in the following fashion. The original formulas are all inverted to give you the original location of the value
	// that finished its journey in the `finalPositionToFind`th position. The offset and multipliers then take the
	// number of desired iterations into account. Then you simply plug in that desired position to the final formula
	//
	// I have heard of using rather large numbers modded by other large primes for use in information obfuscation where
	// the data is useless without both parts (https://www.youtube.com/watch?v=K54ildEW9-Q), though I'm not sure this is
	// 100% the same topic
	n := big.NewInt(length)
	offset, multiplier := big.NewInt(0), big.NewInt(1)
	for _, s := range shuffles {
		switch s.sType {
		case DEAL:
			multiplier.Mul(multiplier, big.NewInt(-1)) // multiplier *= -1
			offset.Add(offset, multiplier)                // offset += multiplier
		case CUT:
			offset.Add(offset, big.NewInt(0).Mul(big.NewInt(int64(s.value)), multiplier)) // offset += cutLength * multiplier
		case DEAL_N:
			multiplier.Mul(multiplier, big.NewInt(0).Exp(big.NewInt(int64(s.value)), big.NewInt(0).Sub(n, big.NewInt(2)), n)) // multiplier *= (dealIncrement ** (length - 2)) % length
		}
	}

	inverseMod := big.NewInt(0).Exp(big.NewInt(0).Sub(big.NewInt(1), multiplier), big.NewInt(0).Sub(n, big.NewInt(2)), n)

	finalIncr := big.NewInt(0).Exp(multiplier, big.NewInt(iterations), n)

	finalOffs := big.NewInt(0).Exp(multiplier, big.NewInt(iterations), n)
	finalOffs.Mul(big.NewInt(-1), finalOffs)
	finalOffs.Add(big.NewInt(1), finalOffs)
	finalOffs.Mul(finalOffs, inverseMod)
	finalOffs.Mul(finalOffs, offset)

	finalFormula := big.NewInt(0).Mul(big.NewInt(finalPositionToFind), finalIncr)
	finalFormula.Add(finalFormula, finalOffs)
	return finalFormula.Mod(finalFormula, n).Int64()
}



