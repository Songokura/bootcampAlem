package bootcamp

func ConvertNbrBase(n int, base string) string {
	if n < 0 || len(base) < 2 || !isUnique(base) {
		return "" // Return empty string for invalid inputs
	}

	baseMap := make(map[byte]int)
	for i := 0; i < len(base); i++ {
		baseMap[base[i]] = i
	}

	var result string
	for n > 0 {
		digit := n % len(base)
		result = string(base[digit]) + result
		n /= len(base)
	}

	return result
}

// isUnique checks if all characters in s are unique
func isUnique(s string) bool {
	seen := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if seen[s[i]] {
			return false
		}
		seen[s[i]] = true
	}
	return true
}

// func main() {
// 	result := ConvertNbrBase(1465, "0123456789")
// 	fmt.Println(result) // 1465

// 	result = ConvertNbrBase(1465, "01")
// 	fmt.Println(result) // 10110111001

// 	result = ConvertNbrBase(1465, "01234567")
// 	fmt.Println(result) // 2671

// 	result = ConvertNbrBase(1465, "0123456789ABCDEF")
// 	fmt.Println(result) // 5B9

// 	result = ConvertNbrBase(1465, "00")
// 	fmt.Println(result) //
// }
