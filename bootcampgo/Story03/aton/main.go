package main

import (
	"bootcamp"
	"fmt"

	"github.com/alem-platform/ap"
)

func main() {
	var N int

	fmt.Scanf("%d", &N)
	arr := make([]rune, N)

	for i := 0; i < N; i++ {
		var temp rune
		fmt.Scanf("%c", &temp)
		arr[i] = temp
	}

	for j := 0; j < N; j++ {
		if 0 < j && j < N {
			ap.PutRune(' ')
		}
		if 0 <= int(arr[j]) && int(arr[j]) <= 127 {
			bootcamp.PutNumber(int(arr[j]))
		} else {
			ap.PutRune(arr[j])
		}
	}
}
