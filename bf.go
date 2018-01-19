package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var file = flag.String("f", "", "file to execute")
	var tapeSize = flag.Int("m", 0, "amount of memory cells. 0 = unlimited")
	// If it's 0 make a dynamic slice
	if *tapeSize == 0 {
		var tape []uint32
	}
	if *tapeSize < 0 {
		fmt.Println("# of memory cells cannot be 0!")
		os.Exit(1)
	}
	if *tapeSize > 0 {
		var tape = make([]uint32, *tapeSize, *tapeSize)
	}
	// The tape has as many cells as can be adressed by an int32
	var iPointer int
	// Read in the program
	if *file == "" {
		fmt.Println("Usage: bf -f <file>")
		os.Exit(1)
	}
	fDesc, err := os.Open(*file)
	if err != nil {
		fmt.Printf("Error while reading file: %v\n", err)
		os.Exit(1)
	}
	defer fDesc.Close()
	programB, err := ioutil.ReadAll(fDesc)
	program := []rune(string(programB))
	if err != nil {
		panic(err)
	}
	// Get maps matching opening and closing brackets
	openCloseMap, closeOpenMap := getBrackPairs(program)

	// Create a bufio reader for stdin
	rd := bufio.NewReader(os.Stdin)
	// Just serially read each instruction
	dPointer := 0
	for iPointer < len(program)-1 {
		token := program[iPointer]
		switch token {
		// TODO: prevent overflow
		case '+':
			tape[dPointer] = tape[dPointer] + 1
			iPointer++
		case '-':
			tape[dPointer] = tape[dPointer] - 1
			iPointer++
		case '>':
			dPointer++
			iPointer++
		case '<':
			dPointer--
			iPointer++
		case '.':
			fmt.Print(string(rune(tape[dPointer])))
			iPointer++
		case ',':
			inStr, err := rd.ReadString('\n')
			if err != nil {
				fmt.Printf("Failed to read from stdin: %v\n", err)
				os.Exit(1)
			}
			inRuneSlice := []rune(inStr)
			if len(inRuneSlice) != 1 {
				fmt.Printf("You can only input 1 character at a time!")
			} else {
				inRune := inRuneSlice[0]
				tape[dPointer] = uint32(inRune)
				if err != nil {
					fmt.Printf("Failed to read from stdin: %v\n", err)
					os.Exit(1)
				}
			}
			iPointer++
		case '[':
			//
			if tape[dPointer] == 0 {
				// Look up what "]" to jump to
				iPointer = openCloseMap[iPointer] + 1

			} else {
				iPointer++
			}
		case ']':
			if tape[dPointer] != 0 {
				// Look up what "[" to jump to
				iPointer = closeOpenMap[iPointer] + 1

			} else {
				iPointer++
			}
		default:
			// Unknown token, do nothing
		}
	}
}

// Parse the program and return a map with [index_of_opening_bracket]index_of_closing_bracket
func getBrackPairs(program []rune) (map[int]int, map[int]int) {
	var openBracks []int
	var closeBracks []int
	// TODO: Paralellize
	// search for "[" from the beginning
	for i, token := range program {
		if token == '[' {
			openBracks = append(openBracks, i)
		}
	}
	// Now search for "]" from the end
	for i := len(program) - 1; i >= 0; i-- {
		if program[i] == ']' {
			closeBracks = append(closeBracks, i)
		}
	}

	if len(openBracks) != len(closeBracks) {
		// TODO: Print index
		fmt.Print("Missing opening or closing bracket!")
		os.Exit(1)
	}
	openCloseMap := make(map[int]int)
	closeOpenMap := make(map[int]int)
	// Now glue them together into a map
	for i := 0; i < len(openBracks); i++ {
		openCloseMap[openBracks[i]] = closeBracks[i]
	}
	// And again in reverse, as go doesn't allow looking up a key by value.
	for i := 0; i < len(closeBracks); i++ {
		closeOpenMap[closeBracks[i]] = openBracks[i]
	}
	return openCloseMap, closeOpenMap
}
