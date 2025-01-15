package bootcamp

func TestNthRune(fn func(s string, n int) rune) bool {
	testCases := []struct {
		s        string
		n        int
		expected rune
	}{
		{"hello", 1, 'h'},
		{"hello", 3, 'l'},
		{"hello", 6, -1},
		{"hello", 0, -1},
		{"", 1, -1},
		{"world", 5, 'd'},
	}

	for _, tc := range testCases {
		result := fn(tc.s, tc.n)
		if result != tc.expected {
			return false
		}
	}

	return true
}

func NthRune(s string, n int) rune {
	if n <= 0 || n > len(s) {
		return -1
	}
	runes := []rune(s)
	return runes[n-1]
}

func ZeroRune(s string, n int) rune {
	return '0'
}

func FirstRune(s string, n int) rune {
	if len(s) > 0 {
		return rune(s[0])
	}
	return -1
}

// func main() {
// 	fmt.Println(TestNthRune(NthRune))   // true
// 	fmt.Println(TestNthRune(ZeroRune))  // false
// 	fmt.Println(TestNthRune(FirstRune)) // false
// }
