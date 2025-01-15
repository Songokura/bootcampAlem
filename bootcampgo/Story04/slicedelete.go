package bootcamp

// import "fmt"

func SliceDelete(arr *[]int, low, high int) bool {
	if low < 0 || low >= len(*arr) || high <= low || high > len(*arr) {
		return false
	}

	numToDelete := high - low

	copy((*arr)[low:], (*arr)[high:])

	*arr = (*arr)[:len(*arr)-numToDelete]

	return true
}

// func main() {
// 	arr := []int{0, 1, 2, 3, 4, 5}

// 	fmt.Println(arr)                     // [0, 1, 2, 3, 4, 5]
// 	fmt.Println(SliceDelete(&arr, 1, 3)) // true
// 	fmt.Println(arr)                     // [0, 3, 4, 5]

// 	fmt.Println(SliceDelete(&arr, 3, 3)) // false
// 	fmt.Println(arr)                     // [0, 3, 4, 5]

// 	fmt.Println(SliceDelete(&arr, 5, 3)) // false
// 	fmt.Println(arr)                     // [0, 3, 4, 5]

// 	fmt.Println(SliceDelete(&arr, -10, 5)) // false
// 	fmt.Println(arr)                       // [0, 3, 4, 5]

// 	fmt.Println(SliceDelete(&arr, 0, 4)) // true
// 	fmt.Println(arr)                     // []
// }
