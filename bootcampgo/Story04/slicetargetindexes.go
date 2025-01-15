package bootcamp

// import "fmt"

func SliceTargetIndexes(arr []int, target int) []int {
	var indexes []int

	for i, val := range arr {
		if val == target {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

// func main() {
// 	arr := []int{2, 3, 2, 5, 6, 7, 8, 2, 10}
// 	fmt.Println(SliceTargetIndexes(arr, 2)) // [0, 2, 7]
// 	fmt.Println(SliceTargetIndexes(arr, 1)) // []
// }
