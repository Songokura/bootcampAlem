package main

import "github.com/alem-platform/ap"

func num_print() {
	for _, v := range "9876543210" {
		ap.PutRune(v)
	}

	ap.PutRune('\n')
}

func main() {
	num_print()
}
