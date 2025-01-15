package bootcamp

import "github.com/alem-platform/ap"

func PrintAlphabetLn() {
	for _, v := range "abcdefghijklmnopqrstuvwxyz" {
		ap.PutRune(v)
	}

	ap.PutRune('\n')
}

func main() {
	PrintAlphabetLn()
}
