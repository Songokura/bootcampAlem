package main

import (
	"bootcamp"
	"fmt"

	"github.com/alem-platform/ap"
)

func main() {
	const maxSize = 20
	var N, num int
	var array [maxSize]int

	fmt.Scanf("%d", &N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &num)
		array[i] = num
	}

	max := array[0]
	min := array[0]
	for i := 1; i < N; i++ {
		if array[i] > max {
			max = array[i]
		}
		if array[i] < min {
			min = array[i]
		}
	}
	bootcamp.PutNumber(max)
	ap.PutRune(' ')
	bootcamp.PutNumber(min)
}
