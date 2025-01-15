package util

import "github.com/alem-platform/ap"

func PrintWall(r rune) {
	ap.PutRune('|')
	for i := 0; i < 7; i++ { // The loop prints the rune 7 times
		ap.PutRune(r)
	}
}

func PrintSpaces(n int, r rune) {
	for i := 0; i < n; i++ { // The loop prints the rune n times
		ap.PutRune(r)
	}
}

func PutRune(r rune) {
	ap.PutRune(r)
}

func HeaderPrint(column int) {
	for i := 0; i < column*8; i++ { // The loop prints the _ column*8 times
		if i != column*8-1 {
			PutRune('_')
		}
	}
}

func PutString(s string) {
	for _, r := range s { // The loop prints the every sign of the string
		ap.PutRune(r)
	}
}
