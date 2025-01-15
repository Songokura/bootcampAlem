package bootcamp

// import "fmt"

func SliceBatch(slice []int, size int) [][]int {
	var batches [][]int

	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		batches = append(batches, slice[i:end])
	}
	return batches
}

// func main() {
// 	arr := []int{1, 2, 3, 4, 5, 6, 7}
// 	batch := SliceBatch(arr, 2)
// 	for _, v := range batch {
// 		fmt.Println(v)
// 	}
// 	// Output:
// 	// [1, 2]
// 	// [3, 4]
// 	// [5, 6]
// 	// [7]
// }
