package bootcamp

// import "fmt"

func isOpeningBracket(b byte) bool {
	return b == '(' || b == '[' || b == '{'
}

func isClosingBracket(b byte) bool {
	return b == ')' || b == ']' || b == '}'
}

func isMatchingPair(open, close byte) bool {
	switch open {
	case '(':
		return close == ')'
	case '[':
		return close == ']'
	case '{':
		return close == '}'
	default:
		return false
	}
}

func Brackets(s string) bool {
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if isOpeningBracket(s[i]) {
			stack = append(stack, s[i])
		} else if isClosingBracket(s[i]) {
			if len(stack) == 0 || !isMatchingPair(stack[len(stack)-1], s[i]) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// func main() {
// 	fmt.Println(Brackets("salem (student) {test} [example]")) // true
// 	fmt.Println(Brackets("test (salem [student)"))            // false
// 	fmt.Println(Brackets("?.fas<. {[...]} (())"))             // true
// 	fmt.Println(Brackets("({[)}]"))                           // false
// }
