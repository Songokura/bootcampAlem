package bootcamp

import "github.com/alem-platform/ap"

func PutNumber(n int) {
	if n < 0 {
		ap.PutRune('-')
		n = -n
	}
	if n == 0 {
		ap.PutRune('0')
		return
	}

	var digits []int

	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	for i := len(digits) - 1; i >= 0; i-- {
		PutDigit(digits[i])
	}
}
