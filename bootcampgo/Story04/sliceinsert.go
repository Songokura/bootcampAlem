package bootcamp

import "fmt"

// import "fmt"

func SliceInsert(arr *[]int, idx int, values ...int) bool {
	if idx < 0 || idx > len(*arr) {
		return false
	}

	if idx == len(*arr) {
		*arr = append(*arr, values...)
		return true
	}

	newLen := len(*arr) + len(values)
	fmt.Println(len(*arr))
	fmt.Println(len(values))
	fmt.Println(cap(*arr))
	if newLen > cap(*arr) {
		newArr := make([]int, newLen)

		copy(newArr[:idx], (*arr)[:idx]) // copy((newarr(2)[:0]), (*arr(1))[:0]

		copy(newArr[idx:], values) // copy((newarr[0:], 10))

		copy(newArr[idx+len(values):], (*arr)[idx:]) // copy(newarr[idx(0)+len(values)(1):], )

		*arr = newArr
	} else {
		*arr = (*arr)[:newLen]
		copy((*arr)[idx+len(values):], (*arr)[idx:])
		copy((*arr)[idx:], values)
	}

	return true
}

// func main() {
// 	arr := []int{1}
// 	fmt.Println(arr) // [1]

// 	fmt.Println(SliceInsert(&arr, 0, 10))
// 	fmt.Println(arr) // [10, 1]

// 	fmt.Println(SliceInsert(&arr, len(arr), 20))
// 	fmt.Println(arr) // [10, 1, 20]

// 	fmt.Println(SliceInsert(&arr, 2, 3)) // true
// 	fmt.Println(arr)                     // [10, 1, 3, 20]

// 	fmt.Println(SliceInsert(&arr, -1, 100)) // false
// 	fmt.Println(arr)                        // [10, 1, 3, 20]
// }
