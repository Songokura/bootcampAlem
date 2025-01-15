package bootcamp

// import "fmt"

func IsAlpha(s string) bool {
	for _, v := range s {
		if (v < 'A' || v > 'Z') && (v < 'a' || v > 'z') {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsAlpha("HelloWorld"))  // true
// 	fmt.Println(IsAlpha("Hello123"))    // false
// 	fmt.Println(IsAlpha("Hello World")) // false
// }
