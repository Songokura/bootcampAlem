package bootcamp

// import "fmt"

func CountSubstr(s string, substr string) int {
	count := 0
	n := len(s)
	m := len(substr)

	for i := 0; i <= n-m; i++ {
		if s[i:i+m] == substr {
			count++
		}
	}
	return count
}

// func main() {
// 	fmt.Println(CountSubstr("qanagattandyrylmagandyqtarynyzdan", "an"))
// 	fmt.Println(CountSubstr("ab ab ab ab ab", "ab"))
// }
