package bootcamp

func MissingNumber(arr []int) int {
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			// Swap arr[i] with arr[j]
			arr[i], arr[j] = arr[j], arr[i]
			j++
		}
	}

	for i := 0; i < j; i++ {

		index := abs(arr[i]) - 1

		if index < j && arr[index] > 0 {
			arr[index] = -arr[index]
		}
	}

	for i := 0; i < j; i++ {
		if arr[i] > 0 {
			return i + 1
		}
	}

	return j + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// func main() {
// 	fmt.Println(MissingNumber([]int{1, 2, 3, 4, 5}))   // 6
// 	fmt.Println(MissingNumber([]int{1, 2, 4, 5}))      // 3
// 	fmt.Println(MissingNumber([]int{3, 4, 5, 6, 7}))   // 1
// 	fmt.Println(MissingNumber([]int{}))                // 1
// 	fmt.Println(MissingNumber([]int{0, -1, -2, 1, 2})) // 3
// }
