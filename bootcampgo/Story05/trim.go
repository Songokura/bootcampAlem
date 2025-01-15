package bootcamp

// import "fmt"

func Trim(s string) string {
	left := 0

	for left < len(s) && (s[left] == ' ' || s[left] == '\t' || s[left] == '\n' || s[left] == '\r') {
		left++
	}

	right := len(s) - 1
	for right >= 0 && (s[right] == ' ' || s[right] == '\t' || s[right] == '\n' || s[right] == '\r') {
		right--
	}

	if left > right {
		return ""
	}
	return s[left : right+1]
}

// func main() {
// 	fmt.Println(Trim("   Salem student!   "))
// 	fmt.Println(Trim("   Trim spaces   "))
// }
