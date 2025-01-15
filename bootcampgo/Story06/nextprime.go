package bootcamp

// import "fmt"

// func IsPrime(n int) bool {
// 	if n <= 1 {
// 		return false
// 	}

// 	for i := 2; i*i <= n; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}

// 	return true
// }

func NextPrime(n int) int {
	if n <= 1 {
		return 2
	}

	for i := n + 1; ; i++ {
		if IsPrime(i) {
			return i
		}
	}
}

// func main() {
// 	fmt.Println(NextPrime(10)) // 11
// 	fmt.Println(NextPrime(11)) // 13
// 	fmt.Println(NextPrime(-1)) // 2
// }
