package bootcamp

// import "fmt"

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	var result string
	isNeg := false

	if n < 0 {
		isNeg = true
		n = -n
	}

	for n > 0 {
		digit := n % 10
		result = string('0'+digit) + result
		n /= 10
	}

	if isNeg {
		result = "-" + result
	}
	return result
}

// func main() {
// 	fmt.Println(Itoa(123))
// 	fmt.Println(Itoa(-456))
// 	fmt.Println(Itoa(0))
// }
