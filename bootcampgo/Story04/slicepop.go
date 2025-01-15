package bootcamp

// import "fmt"

func SlicePop(arr *[]int) int {
	if len(*arr) == 0 {
		return 0
	}
	lastIndex := len(*arr) - 1
	poppedElement := (*arr)[lastIndex]
	*arr = (*arr)[:lastIndex]
	return poppedElement
}

// func main() {
// 	arr := []int{5, 10, 20}
// 	size := len(arr)

// 	for i := 0; i < size; i++ {
// 		deleted := SlicePop(&arr)
// 		fmt.Println(deleted)
// 	}
// 	// Output:
// 	// 20
// 	// 10
// 	// 5

// 	deleted := SlicePop(&arr)
// 	fmt.Println(deleted) // 0
// }
