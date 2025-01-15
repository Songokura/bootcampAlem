package bootcamp

// import "fmt"

func Repeat(s string, count int) string {
	if count <= 0 {
		return ""
	}

	back := ""
	for i := 0; i < count; i++ {
		back += s
	}
	return back
}

// func main() {
// 	fmt.Println(Repeat("abc", 3)) // "abcabcabc"
// 	fmt.Println(Repeat("123", 2)) // "123123"
// }
