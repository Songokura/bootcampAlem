package bootcamp

// import "fmt"

func DigitsOnly(s string) string {
	var digits []rune

	for _, v := range s {
		if v >= '0' && v <= '9' {
			digits = append(digits, v)
		}
	}
	return string(digits)
}

// func main() {
// 	fmt.Println(DigitsOnly("abc123"))
// 	fmt.Println(DigitsOnly("Sa1em student! 23"))
// 	fmt.Println(DigitsOnly("no digits here!"))
// }
