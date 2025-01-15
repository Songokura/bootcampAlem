package utils

import "github.com/alem-platform/ap"

// This function prints any strings
func Print(a string) {
	ap.PutRune('\n')
	for _, v := range a {
		ap.PutRune(v)
	}
	ap.PutRune('\n')
}
