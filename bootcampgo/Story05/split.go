package bootcamp

// import "fmt"

func Split(s string, sep string) []string {
	var result []string

	if sep == "" {
		// Edge case: if the separator is empty, return each character as a separate string
		for _, char := range s {
			result = append(result, string(char))
		}
		return result
	}

	sepLen := len(sep)
	for {
		index := indexOf(s, sep)
		if index == -1 {
			result = append(result, s)
			break
		}
		result = append(result, s[:index])
		s = s[index+sepLen:]
	}

	return result
}

func indexOf(s, sep string) int {
	n := len(s)
	m := len(sep)

	for i := 0; i <= n-m; i++ {
		if s[i:i+m] == sep {
			return i
		}
	}

	return -1
}

// func main() {
// 	fmt.Println(Split("a,b,c", ","))
// 	fmt.Println(Split("salem-student", "-"))
// 	fmt.Println(Split("salem", ""))
// }
