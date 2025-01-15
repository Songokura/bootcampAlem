package bootcamp

// import "fmt"

func Rot13(s string) string {
	var result string

	for _, v := range s {
		switch {
		case 'a' <= v && v <= 'z':
			result += string('a' + (v-'a'+13)%26)
		case 'A' <= v && v <= 'Z':
			result += string('A' + (v-'A'+13)%26)
		default:
			result += string(v)
		}
	}
	return result
}

// func main() {
// 	fmt.Println(Rot13("salem"))
// 	fmt.Println(Rot13("fnyrz"))
// }
