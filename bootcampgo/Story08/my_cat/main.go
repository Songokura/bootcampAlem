package main

import (
	"os"

	"github.com/alem-platform/ap"
)

// displayText prints a string using ap.PutRune for each character
func Print(text string) {
	for _, ch := range text {
		ap.PutRune(ch)
	}
}

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		text := []rune("usage: my_cat file ...")
		for _, v := range text {
			ap.PutRune(v)
		}
		ap.PutRune('\n')
		return
	}
	for i := 0; i < len(files); i++ {
		b, err := os.ReadFile(files[i])
		if err != nil {
			Print("my_cat: ")
			Print(files[i])
			Print(": No such file or directory\n")
			continue
		}
		Print(string(b))
	}
}
