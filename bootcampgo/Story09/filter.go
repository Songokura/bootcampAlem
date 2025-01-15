package bootcamp

func Filter(arr []int, fn func(int) bool) []int {
	result := []int{}
	for _, v := range arr {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// func main() {
// 	isEven := func(n int) bool {
// 		return n%2 == 0
// 	}

// 	result := Filter([]int{1, 2, 3, 4, 5, 6}, isEven)
// 	fmt.Println(result) // [2 4 6]

// 	isGreaterThanThree := func(n int) bool {
// 		return n > 3
// 	}

// 	result = Filter([]int{1, 2, 3, 4, 5, 6}, isGreaterThanThree)
// 	fmt.Println(result) // [4 5 6]
// }
