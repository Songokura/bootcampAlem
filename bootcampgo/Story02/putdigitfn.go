package bootcamp

import "github.com/alem-platform/ap"

func PutDigit(n int) {
	if n >= 0 && n <= 9 {
		ap.PutRune('0' + rune(n%10))
	}
}

//func main() {
//	PutDigit(21)
//}
