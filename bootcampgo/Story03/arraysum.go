package bootcamp

import "github.com/alem-platform/ap"

func ArraySum(arr [20]int) int {
	sum := 0
	for _, num := range arr {
		sum += num
		ap.PutRune(' ')
	}
	return sum
}
