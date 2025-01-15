package bootcamp

// import "fmt"

// func allUnique(characters string) bool {
// 	seen := make(map[rune]bool)
// 	for _, char := range characters {
// 		if seen[char] {
// 			return false
// 		}
// 		seen[char] = true
// 	}
// 	return true
// }

// Recursive function to generate combinations
func generateCombinations(prefix string, characters []rune, n int, result *[]string) {
	if n == 0 {
		*result = append(*result, prefix)
		return
	}
	for i := 0; i < len(characters); i++ {
		newPrefix := prefix + string(characters[i])
		newChars := append([]rune{}, characters[:i]...)
		newChars = append(newChars, characters[i+1:]...)
		generateCombinations(newPrefix, newChars, n-1, result)
	}
}

// Function to generate all unique n-character combinations
func UniqCombN(characters string, n int) []string {
	// Check if the characters are unique
	if !allUnique(characters) {
		return []string{}
	}

	// Check if it is possible to generate n-character combinations
	if n > len(characters) || n < 1 {
		return []string{}
	}

	var result []string
	chars := []rune(characters)

	// Generate all possible combinations of n distinct characters
	generateCombinations("", chars, n, &result)

	return result
}

// func main() {
// 	fmt.Println(UniqCombN("abc", 1)) // ["a", "b", "c"]
// 	fmt.Println(UniqCombN("abc", 2)) // ["ab", "ac", "ba", "bc", "ca", "cb"]
// 	fmt.Println(UniqCombN("ab", 2))  // ["ab", "ba"]
// 	fmt.Println(UniqCombN("a", 1))   // ["a"]
// 	fmt.Println(UniqCombN("ab", 3))  // []
// 	fmt.Println(UniqCombN("aa", 1))  // []
// }
