package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func main() {
	// Iterate through os.Args to print each argument
	for _, arg := range os.Args[1:] {
		printArgument(arg)
	}
}

// printArgument prints each character of the given argument followed by a newline
func printArgument(arg string) {
	// Print each character of the argument
	for _, char := range arg {
		ap.PutRune(char)
	}
	// Print newline after printing the argument
	ap.PutRune('\n')
}
