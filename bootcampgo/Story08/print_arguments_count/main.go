package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func main() {
	// Get the number of arguments passed to the program
	numArgs := len(os.Args) - 1 // Exclude the program name itself

	// Print the number of arguments digit by digit
	printNumber(numArgs)
	ap.PutRune('\n')
}

// printNumber prints the given number digit by digit
func printNumber(num int) {
	// Handle the case where num is 0 separately
	if num == 0 {
		ap.PutRune('0')
		return
	}

	// Print each digit of the number
	printDigits(num)
}

// printDigits prints the digits of a number recursively
func printDigits(num int) {
	// Base case: if num is 0, stop recursion
	if num == 0 {
		return
	}

	// Recursive call to print remaining digits
	printDigits(num / 10)

	// Print the current digit
	digit := '0' + rune(num%10)
	ap.PutRune(digit)
}
