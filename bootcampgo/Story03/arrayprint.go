package bootcamp

import "github.com/alem-platform/ap"

func PutD(n int) {
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
		PutD(digits[k])
	}
}

func ArrayPrint(arr [20]int) {
	for i, num := range arr {
		conv(num)
		if i < len(arr)-1 {
			ap.PutRune(' ')
		}
	}
}
