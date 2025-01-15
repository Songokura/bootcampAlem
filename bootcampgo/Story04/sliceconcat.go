package bootcamp

// import "fmt"

func SliceConcat(slices ...[]int) []int {
	var concatendatedSlice []int

	for _, v := range slices {
		concatendatedSlice = append(concatendatedSlice, v...)
	}
	return concatendatedSlice
}

// func main() {
// 	result := SliceConcat([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4}, []int{15, 0, 2})
// 	fmt.Println(result) // [1, 2, 3, 4, 5, 1, 2, 3, 4, 15, 0, 2]
// }
