package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

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

func sum(a, b int) {
	sum := a + b
	conv(sum)
	ap.PutRune(' ')
}

func subs(a, b int) {
	subs := a - b
	conv(subs)
	ap.PutRune(' ')
}

func mult(a, b int) {
	mult := a * b
	conv(mult)
	ap.PutRune(' ')
}

func div(a, b int) {
	if b != 0 {
		div := a / b
		conv(div)
	} else {
		ap.PutRune('F')
	}
}

func main() {
	var a, b int

	fmt.Scanf("%d %d", &a, &b)

	sum(a, b)
	subs(a, b)
	mult(a, b)
	div(a, b)
}
