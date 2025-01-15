package main

import "github.com/alem-platform/ap"

func alp_print() {
	for _, v := range "zyxwvutsrqponmlkjihgfedcba" {
		ap.PutRune(v)
	}

	ap.PutRune('\n')
}

func main() {
	alp_print()
}
