package bootcamp

// import "fmt"

func Contains(str string, substr string) bool {
	strLen := len(str)
	substrLen := len(substr)

	if substrLen > strLen {
		return false
	}

	for i := 0; i <= strLen-substrLen; i++ {
		if str[i:i+substrLen] == substr {
			return true
		}
	}

	return false
}

// func main() {
// 	fmt.Println(Contains("hello world", "world"))
// 	fmt.Println(Contains("test", "dest"))
// }
