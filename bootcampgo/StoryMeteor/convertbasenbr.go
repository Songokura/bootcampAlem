package bootcamp

// import "fmt"

func ConvertBaseNbr(n string, base string) int {
	if len(base) < 2 || !IsUnique(base) {
		return -1 // Return -1 for invalid base
	}

	baseMap := make(map[byte]int)
	for i := 0; i < len(base); i++ {
		baseMap[base[i]] = i
	}

	result := 0
	for i := 0; i < len(n); i++ {
		char := n[i]
		if val, ok := baseMap[char]; ok {
			result = result*len(base) + val
		} else {
			return -1 // Return -1 if any character in n is not in base
		}
	}

	return result
}

// isUnique checks if all characters in s are unique
func IsUnique(s string) bool {
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
// 	fmt.Println(ConvertBaseNbr("1465", "0123456789"))      // 1465
// 	fmt.Println(ConvertBaseNbr("10110111001", "01"))       // 1465
// 	fmt.Println(ConvertBaseNbr("2671", "01234567"))        // 1465
// 	fmt.Println(ConvertBaseNbr("5B9", "0123456789ABCDEF")) // 1465
// 	fmt.Println(ConvertBaseNbr("1", "00"))                 // -1
// }
