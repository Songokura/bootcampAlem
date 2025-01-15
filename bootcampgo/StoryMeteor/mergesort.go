package bootcamp

import (
	"github.com/alem-platform/ap"
)

func MergeSort(arr []int) {
	n := len(arr)
	swapped := true

	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapped = true
			}
		}
	}

	printLastLine(arr)
}

func printLastLine(arr []int) {
	for _, num := range arr {
		if num < 0 {
			ap.PutRune('-')
			num = -num
		}

		printDigits(num)

		ap.PutRune(' ')
	}
	ap.PutRune('\n')
}

func printDigits(num int) {
	if num == 0 {
		ap.PutRune('0')
		return
	}

	digits := make([]int, 0)

	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		ap.PutRune(rune('0' + digits[i]))
	}
}

// func main() {
// 	arr := []int{38, 27, 43, 3, 9, 82, 10}
// 	MergeSort(arr)
// 	fmt.Println(arr) // Print newline after outputting all symbols
// }
