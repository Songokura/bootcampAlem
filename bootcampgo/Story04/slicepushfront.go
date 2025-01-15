package bootcamp

// import "fmt"

func SlicePushFront(arr *[]int, values ...int) {
	*arr = append(values, (*arr)...)
}

// func main() {
// 	arr := []int{1}
// 	fmt.Println(arr) // [1]

// 	SlicePushFront(&arr, 10)
// 	fmt.Println(arr) // [10, 1]

// 	SlicePushFront(&arr, 20)
// 	fmt.Println(arr) // [20, 10, 1]

// 	SlicePushFront(&arr, 5, 3)
// 	fmt.Println(arr) // [5, 3, 20, 10, 1]
// }
