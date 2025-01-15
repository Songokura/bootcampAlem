package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func atoi(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}

	var num int
	sign := 1
	i := 0

	// Handle leading whitespaces
	for i < len(s) && s[i] == ' ' {
		i++
	}

	// Handle sign
	if i < len(s) && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	// Convert digits
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		num = num*10 + int(s[i]-'0')
	}

	return num * sign, true
}

func PutDigit(n int) {
	if n >= 0 && n < 10 {
		ap.PutRune('0' + rune(n%10))
	}
}

func conv(i int) {
	if i < 0 {
		ap.PutRune('-')
		i = -i
	}
	if i == 0 {
		ap.PutRune('0')
		return
	}

	var digits []int

	for i > 0 {
		digits = append(digits, i%10)
		i /= 10
	}
	for k := len(digits) - 1; k >= 0; k-- {
		PutDigit(digits[k])
	}
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	arg := os.Args[1]
	num, ok := atoi(arg)
	if !ok || num <= 0 {
		return
	}

	for num != 1 {
		conv(num)
		ap.PutRune('\n')
		if num%2 == 0 {
			num = num / 2
		} else {
			num = num*3 + 1
		}
	}
	conv(num)
	ap.PutRune('\n') // Print the final 1
}
