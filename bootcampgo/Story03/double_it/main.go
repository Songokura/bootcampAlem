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

func main() {
	var L int

	fmt.Scanf("%d", &L)

	for i := 0; i < L; i++ {
		var num int
		fmt.Scanf("%d", &num)
		conv(num * 2)
		ap.PutRune(' ')
	}
}
