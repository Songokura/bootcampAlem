package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func Print(text string) {
	for _, ch := range text {
		ap.PutRune(ch)
	}
}

func main() {
	cw, cl, cb := false, false, false
	arg := os.Args[1]
	prm := ""
	filename := []string{}
	for i := 0; i < len(arg); i++ {
		if arg[0] == '-' {
			prm = os.Args[1]
			filename = os.Args[2:]
		} else {
			filename = os.Args[1:]
		}
	}
	for i := 0; i < len(filename); i++ {
		b, err := os.ReadFile(filename[i])
		if err != nil {
			Print("my_wc: ")
			Print(filename[i])
			Print(": No such file or directory\n")
			return
		}
		words := string(b)
		for _, value := range prm {
			if value == 'l' {
				cl = true
			}
			if value == 'w' {
				cw = true
			}
			if value == 'c' {
				cb = true
			}
		}
		if !cl && !cw && !cb {
			Print(Itoa(countnewLine(words)))
			Print(" ")
			Print(Itoa(countWords(words)))
			Print(" ")
			Print(Itoa(countbytes(words)))
			Print(" ")
		}
		if cl {
			Print(Itoa(countnewLine(words)))
			Print(" ")
		}
		if cw {
			Print(Itoa(countWords(words)))
			Print(" ")
		}
		if cb {
			Print(Itoa(countbytes(words)))
			Print(" ")
		}
		Print(filename[i])
		Print("\n")
	}
}

func countnewLine(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 10 {
			count++
		}
	}
	return count
}

func countWords(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			count++
			for k := 0; k < len(s); k++ {
				if k+i == len(s)-1 {
					i = i + k
					break
				}
				if s[i+k] == '\n' {
					i = i + k
					break
				}
				if s[i+k] == 32 {
					i = i + k
					break
				}
			}
		}
	}
	return count
}

func countbytes(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 0 && s[i] <= 255 {
			count++
		}
	}
	return count
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	var result string
	isNeg := false

	if n < 0 {
		isNeg = true
		n = -n
	}

	for n > 0 {
		digit := n % 10
		result = string('0'+digit) + result
		n /= 10
	}

	if isNeg {
		result = "-" + result
	}
	return result
}
