package bootcamp

func IsSorted(arr []int, fn func(int, int) bool) bool {
	for i := 1; i < len(arr); i++ {
		if !fn(arr[i-1], arr[i]) {
			return false
		}
	}
	return true
}

// func main() {
// 	ascending := func(a, b int) bool {
// 		return a <= b
// 	}

// 	descending := func(a, b int) bool {
// 		return a >= b
// 	}

// 	result := IsSorted([]int{1, 2, 3, 4, 5}, ascending)
// 	fmt.Println(result) // true

// 	result = IsSorted([]int{5, 4, 3, 2, 1}, ascending)
// 	fmt.Println(result) // false

// 	result = IsSorted([]int{5, 4, 3, 2, 1}, descending)
// 	fmt.Println(result) // true

// 	result = IsSorted([]int{1, 3, 2, 4, 5}, descending)
// 	fmt.Println(result) // false
// }
