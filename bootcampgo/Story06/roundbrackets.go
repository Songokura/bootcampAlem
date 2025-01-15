package bootcamp

// import "fmt"

func RoundBrackets(s string) bool {
	balance := 0

	if len(s) == 0 || s[0] == ')' {
		return false
	}

	for _, v := range s {
		if v == '(' {
			balance++
		} else if v == ')' {
			balance--
		}
	}
	return balance == 0
}

// func main() {
// 	fmt.Println(RoundBrackets("()()()"))   // true
// 	fmt.Println(RoundBrackets("(salem)"))  // true
// 	fmt.Println(RoundBrackets(")(1)(f)(")) // false
// 	fmt.Println(RoundBrackets("))(()"))    // false
// 	fmt.Println(RoundBrackets(""))         // false
// }
