package bootcamp

func RomanToInt(s string) int {
	romanValues := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	if len(s) == 0 {
		return 0
	}

	total := 0
	prev := 0

	// Traverse the string from right to left
	for i := len(s) - 1; i >= 0; i-- {
		current := romanValues[s[i]]

		// Check if we need to add or subtract
		if current >= prev {
			total += current
		} else {
			total -= current
		}

		prev = current
	}

	return total
}

// func main() {
// 	fmt.Println(RomanToInt("III"))     // 3
// 	fmt.Println(RomanToInt("IV"))      // 4
// 	fmt.Println(RomanToInt("IX"))      // 9
// 	fmt.Println(RomanToInt("LVIII"))   // 58
// 	fmt.Println(RomanToInt("MCMXCIV")) // 1994
// 	fmt.Println(RomanToInt(""))        // 0
// 	fmt.Println(RomanToInt("salem"))   // 0
// }
