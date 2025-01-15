package bootcamp

// import "fmt"

func SliceSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

// func main() {
// 	arr := []int{10, 90, 20, 5, 12, 3, 0}
// 	SliceSort(arr)
// 	fmt.Println(arr) // [0, 3, 5, 10, 12, 20, 90]
// }
