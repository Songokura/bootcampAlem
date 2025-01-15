package bootcamp

// import "fmt"

func SliceCopy(dst, src []int) {
	copy(dst, src)
}

// func main() {
// 	src := []int{10, 20, 13, 5, 12, 31}
// 	dst := make([]int, 4)

// 	SliceCopy(dst, src)

// 	fmt.Println(src, dst) // [10, 20, 13, 5, 12, 31] [10, 20, 13, 5]

// 	src[0] = 0

// 	fmt.Println(src, dst) // [0, 20, 13, 5, 12, 31] [10, 20, 13, 5]
// }
