package bootcamp

import "github.com/alem-platform/ap"

func PrintBits(n byte) {
	for i := 7; i >= 0; i-- {
		bit := (n >> uint(i)) & 1
		if bit == 1 {
			ap.PutRune('1')
		} else {
			ap.PutRune('0')
		}
	}
	ap.PutRune('\n')
}

// func main() {
// 	PrintBits(5)
// 	// 00000101
// 	PrintBits(255)
// 	// 11111111
// 	PrintBits(0)
// 	// 00000000
// }
