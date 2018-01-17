package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// The tape has as many cells as can be adressed by an int32
	var tape [2147483648]uint32
	var pointer uint32
	var bracketCount int
	// Read in the program
	file := os.Args[1]
	programB, err := ioutil.ReadAll(file)
	program := rune(string(programB))
	if err != nil {
		panic(err)
	}

	// Create a bufio reader for stdin
	rd := bufio.NewReader(os.Stdin)
	// Just serially read each instruction
	for i, token := range program {
		switch token {
		// TODO: prevent overflow
		case "+":
			tape[pointer] = tape[pointer] + 1
		case "-":
			tape[pointer] = tape[pointer] - 1
		case ">":
			pointer++
		case "<":
			pointer--
		case ".":
			fmt.Printf(tape[pointer])
		case ",":
			rd.ReadString('\n')
			// TODO
			// TODO: Validate that there are no loose "]" without a match
		case "[":
			// Keep track of the number of brackets passed so we can match.
			bracketCount++
			if tape[pointer] == 0 {
				// Read all tokens until the matching "]" is found

			}
		case "]":
			// TODO
		default:
			// Invalid token, do nothing
		}
	}
}

func getBrackPairs(program string) (brackPairs map[int]int) {
	// search for "["
	var leftBracks []int
	var rightBracks []int
	// TODO: Paralellize
	for i, token := range program {
		if token == "[" {
			append(leftBracks, i)
		}

	}
	// Now search for "]"
	for i := len(program) - 1; i >= 0; i++ {
		if program[i] == "]" {
			append(rightBracks, i)
		}
	}

	// Now glue them together into a map
	for i := 0; i < len(leftBracks); i++ {

	}
}
