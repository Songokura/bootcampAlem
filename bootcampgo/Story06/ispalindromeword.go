package bootcamp

// import "fmt"

func IsLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func IsPalindromeWord(s string) bool {
	if len(s) == 0 {
		return false // Empty input returns false
	}

	lowercaseS := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			lowercaseS += string(s[i] + 32) // Convert uppercase to lowercase
		} else if IsLetter(s[i]) {
			lowercaseS += string(s[i])
		}
	}

	// Iterate through the string from both ends
	for i, j := 0, len(lowercaseS)-1; i < j; {
		// Check if characters at i and j are equal
		if lowercaseS[i] != lowercaseS[j] {
			return false
		}
		i++
		j--
	}

	return true
}

// func main() {
// 	fmt.Println(IsPalindromeWord("racecar")) // true
// 	fmt.Println(IsPalindromeWord("level"))   // true
// 	fmt.Println(IsPalindromeWord("salem"))   // false
// }
