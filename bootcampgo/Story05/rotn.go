package bootcamp

// import "fmt"

func RotN(s string, n int) string {
	var result string

	for _, v := range s {
		switch {
		case 'a' <= v && v <= 'z':
			result += string('a' + (v-'a'+rune(n))%26)
		case 'A' <= v && v <= 'Z':
			result += string('A' + (v-'A'+rune(n))%26)
		default:
			result += string(v)
		}
	}
	return result
}

// func main() {
// 	fmt.Println(RotN("salem", 1))
// 	fmt.Println(RotN("abc", 13))
// }
