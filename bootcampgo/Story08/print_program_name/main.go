package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func main() {
	programPath := os.Args[0]

	// Find the last index of '/' in the path
	lastSlashIndex := 0
	for i := len(programPath) - 1; i >= 0; i-- {
		if programPath[i] == '/' {
			lastSlashIndex = i
			break
		}
	}

	// Extract the program name from the path
	programName := programPath[lastSlashIndex+1:]

	// Print the program name using ap.PutRune
	for _, char := range programName {
		ap.PutRune(char)
	}
	ap.PutRune('\n')
}
