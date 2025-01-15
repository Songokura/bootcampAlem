package main

import "github.com/alem-platform/ap"

func alp_print() {
	for _, v := range "abcdefghijklmnopqrstuvwxyz" {
		ap.PutRune(v)
	}

	ap.PutRune('\n')
}

func main() {
	alp_print()
}
