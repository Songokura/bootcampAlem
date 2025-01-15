package bootcamp

// import "fmt"

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// func main() {
// 	fmt.Println(IsPrime(11))  // true
// 	fmt.Println(IsPrime(4))   // false
// 	fmt.Println(IsPrime(-11)) // false
// }
