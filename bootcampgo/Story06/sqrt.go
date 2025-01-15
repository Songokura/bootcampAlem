package bootcamp

// package main

// import "fmt"

func Sqrt(x int) int {
	if x < 0 {
		return -1
	}

	sqrt := 0
	for i := 0; i*i <= x; i++ {
		if i*i == x {
			sqrt = i
			break
		}
	}
	if sqrt*sqrt == x {
		return sqrt
	}

	return -1
}

// func main() {
// 	fmt.Println(Sqrt(16)) // 4
// 	fmt.Println(Sqrt(1))  // 1
// 	fmt.Println(Sqrt(3))  // 0
// 	fmt.Println(Sqrt(-4)) // -1
// }
