package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func executeBrainfuck(code string) {
	memory := make([]byte, 30000) // Brainfuck memory tape
	ptr := 0                      // Pointer to current memory cell
	codePtr := 0                  // Pointer to current position in code

	output := make([]rune, 0) // To accumulate output for the '.' command

	for codePtr < len(code) {
		switch code[codePtr] {
		case '>':
			ptr++
		case '<':
			ptr--
		case '+':
			memory[ptr]++
		case '-':
			memory[ptr]--
		case '.':
			output = append(output, rune(memory[ptr]))
		case '[':
			if memory[ptr] == 0 {
				// Find matching ']'
				loopLevel := 1
				for loopLevel > 0 {
					codePtr++
					if code[codePtr] == '[' {
						loopLevel++
					} else if code[codePtr] == ']' {
						loopLevel--
					}
				}
			}
		case ']':
			if memory[ptr] != 0 {
				// Find matching '['
				loopLevel := 1
				for loopLevel > 0 {
					codePtr--
					if code[codePtr] == ']' {
						loopLevel++
					} else if code[codePtr] == '[' {
						loopLevel--
					}
				}
				codePtr-- // Adjust to point to the '['
			}
		}
		codePtr++
	}

	// Print accumulated output
	for _, r := range output {
		ap.PutRune(r)
	}
}

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	code := os.Args[1]
	executeBrainfuck(code)
}
