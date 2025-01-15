package bootcamp

// import "fmt"

func Replace(s string, old string, new string) string {
	var result string

	for i := 0; i < len(s); {
		if len(s)-i >= len(old) && s[i:i+len(old)] == old {
			result += new
			i += len(old)
		} else {
			result += string(s[i])
			i++
		}
	}

	return result
}

// func main() {
// 	fmt.Println(Replace("Hello student!", "student", "Mate"))
// 	fmt.Println(Replace("banana", "a", "o"))
// }
