package bootcamp

// import "fmt"

func ToLower(s string) string {
	var result string
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			result += string(v - 'A' + 'a')
		} else {
			result += string(v)
		}
	}
	return result
}

// func main() {
// 	fmt.Println(ToLower("SALEM "))        // "salem "
// 	fmt.Println(ToLower("Salem Student")) // "salem student"
// 	fmt.Println(ToLower("S4LEm"))         // "s4lem"
// }
