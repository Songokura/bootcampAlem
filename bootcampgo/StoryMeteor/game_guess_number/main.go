package main

import (
	"fmt"
	"math/rand"

	"github.com/alem-platform/ap"
)

func PrintW(s string) {
	for _, char := range s {
		ap.PutRune(char)
	}
}

func GuessingGame() {
	rand.Seed(rand.Int63())
	secretNumber := rand.Intn(100) + 1
	var guess int

	for {
		PrintW("Guess number: ")

		_, err := fmt.Scanf("%d\n", &guess)
		if err != nil {
			PrintW("Invalid input. Please enter a valid number.\n")
			continue
		}

		if guess < 1 || guess > 100 {
			continue
		}

		if guess < secretNumber {
			PrintW("Higher\n")
		} else if guess > secretNumber {
			PrintW("Lower\n")
		} else {
			PrintW("Match, you win!\n")
			break
		}
	}
}

func main() {
	GuessingGame()
}
