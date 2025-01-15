package bootcamp

// import (
// 	"fmt"
// )

func Capitalize(s string) string {
	var result string
	capitalizeNext := true
	for _, v := range s {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || (v >= '!' && v <= '@') || (v >= '[' && v <= '`') || (v >= '{' && v <= '~') {
			if capitalizeNext {
				if v >= 'a' && v <= 'z' {
					result += string(v - 'a' + 'A')
				} else {
					result += string(v)
				}
				capitalizeNext = false
			} else {
				if v >= 'A' && v <= 'Z' {
					result += string(v + 'a' - 'A')
				} else {
					result += string(v)
				}
			}
		} else {
			result += string(v)
			capitalizeNext = true
		}
	}
	return result
}

// func main() {
// 	fmt.Println(Capitalize("f[[?s"))
// 	fmt.Println(Capitalize("fjasdnkmsfmsa121313###$$#@#!@#@@, uMER*(*@E&R? &&&&&Dsdasdasdasdsadsa 890gggfE"))
// }
