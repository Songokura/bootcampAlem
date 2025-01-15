package bootcamp

func MergeSortedArrays(arr1, arr2 []int) []int {
	i, j := 0, 0
	merged := make([]int, 0, len(arr1)+len(arr2))

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			merged = append(merged, arr1[i])
			i++
		} else {
			merged = append(merged, arr2[j])
			j++
		}
	}

	// Append any remaining elements from arr1
	for i < len(arr1) {
		merged = append(merged, arr1[i])
		i++
	}

	// Append any remaining elements from arr2
	for j < len(arr2) {
		merged = append(merged, arr2[j])
		j++
	}

	return merged
}

// func main() {
// 	arr1 := []int{1, 3, 5, 7}
// 	arr2 := []int{2, 4, 6, 8}
// 	result := MergeSortedArrays(arr1, arr2)
// 	fmt.Println(result) // [1 2 3 4 5 6 7 8]

// 	arr1 = []int{0, 2, 4}
// 	arr2 = []int{1, 3, 5}
// 	result = MergeSortedArrays(arr1, arr2)
// 	fmt.Println(result) // [0 1 2 3 4 5]

// 	arr1 = []int{10, 20, 30}
// 	arr2 = []int{5, 15, 25, 35}
// 	result = MergeSortedArrays(arr1, arr2)
// 	fmt.Println(result) // [5 10 15 20 25 30 35]
// }
