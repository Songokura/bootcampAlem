package bootcamp

// import "fmt"

func PowRecursion(x int, power int) int {
	if power < 0 {
		return -1
	}
	if power == 0 {
		return 1
	}
	return x * PowRecursion(x, power-1)
}

// func main() {
// 	fmt.Println(PowRecursion(2, -1)) // -1
// 	fmt.Println(PowRecursion(2, 0))  // 1
// 	fmt.Println(PowRecursion(2, 1))  // 2
// 	fmt.Println(PowRecursion(2, 2))  // 4
// 	fmt.Println(PowRecursion(2, 3))  // 8
// 	fmt.Println(PowRecursion(2, 4))  // 16
// }
