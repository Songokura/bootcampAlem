package bootcamp

// import "fmt"

func allUnique(characters string) bool {
	seen := make(map[rune]bool)
	for _, char := range characters {
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}

func UniqComb2(characters string) []string {
	// Check if the characters are unique
	if !allUnique(characters) {
		return []string{}
	}

	// Check if it is possible to generate two-character combinations
	if len(characters) < 2 {
		return []string{}
	}

	var result []string
	chars := []rune(characters)

	// Generate all possible combinations of two distinct characters
	for i := 0; i < len(chars); i++ {
		for j := 0; j < len(chars); j++ {
			if i != j {
				result = append(result, string(chars[i])+string(chars[j]))
			}
		}
	}

	return result
}

// func main() {
// 	fmt.Println(UniqComb2("abc")) // ["ab", "ac", "ba", "bc", "ca", "cb"]
// 	fmt.Println(UniqComb2("ab"))  // ["ab", "ba"]
// 	fmt.Println(UniqComb2("a"))   // []
// 	fmt.Println(UniqComb2("aba")) // []
// }
