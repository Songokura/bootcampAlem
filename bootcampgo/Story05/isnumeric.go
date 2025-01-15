package bootcamp

// import "fmt"

func IsNumeric(s string) bool {
	plusCount, minusCount := 0, 0
	prevWasSign := false

	for _, char := range s {
		if char >= '0' && char <= '9' {
			prevWasSign = false
		} else if char == '-' || char == '+' {
			if prevWasSign {
				return false // More than one sign together
			}
			prevWasSign = true
			if char == '-' {
				minusCount++
			} else {
				plusCount++
			}
		} else {
			return false // Non-numeric character
		}
	}
	return !(plusCount > 1 || minusCount > 1)
}

// func main() {
// 	fmt.Println(IsNumeric("123"))     // true
// 	fmt.Println(IsNumeric("-123"))    // true
// 	fmt.Println(IsNumeric("+123"))    // true
// 	fmt.Println(IsNumeric("+-123"))   // false
// 	fmt.Println(IsNumeric(" 123"))    // false
// 	fmt.Println(IsNumeric("123 abc")) // false
// 	fmt.Println(IsNumeric("123abc"))  // false
// 	fmt.Println(IsNumeric("ab"))      // false
// }
