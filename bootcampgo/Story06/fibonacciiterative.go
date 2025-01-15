package bootcamp

// import "fmt"

func FibonacciIterative(n int) int {
	if n < 0 {
		return -1
	}
	if n <= 1 {
		return n
	}
	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}

// func main() {
// 	fmt.Println(FibonacciIterative(-1)) // -1
// 	fmt.Println(FibonacciIterative(0))  // 0
// 	fmt.Println(FibonacciIterative(1))  // 1
// 	fmt.Println(FibonacciIterative(2))  // 1
// 	fmt.Println(FibonacciIterative(3))  // 2
// 	fmt.Println(FibonacciIterative(4))  // 3
// 	fmt.Println(FibonacciIterative(5))  // 5
// 	fmt.Println(FibonacciIterative(6))  // 8
// }
