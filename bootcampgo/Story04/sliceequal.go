package bootcamp

// import "fmt"

func SliceEqual(arr1, arr2 []int) bool {
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

// func main() {
// 	arr1 := []int{77, 69, 76, 65}
// 	arr2 := []int{77, 69, 76, 65}
// 	fmt.Println(SliceEqual(arr1, arr2))          // true
// 	fmt.Println(SliceEqual(arr1, []int{77, 78})) // false
// }
