package bootcamp

// import "fmt"

func ToUpper(s string) string {
	var result string

	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			result += string(v - 'a' + 'A')
		} else {
			result += string(v)
		}
	}
	return result
}

// func main() {
// 	fmt.Println(ToUpper("salem "))        // "SALEM "
// 	fmt.Println(ToUpper("Salem Student")) // "SALEM STUDENT"
// 	fmt.Println(ToUpper("S4LEm"))         // "S4LEM"
// }
