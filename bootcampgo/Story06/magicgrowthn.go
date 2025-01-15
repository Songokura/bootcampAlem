package bootcamp

// import "fmt"

func MagicGrowthN(n int) []string {
	if n < 1 || n > 10 {
		return []string{}
	}

	var result []string
	digits := "0123456789"

	var generateCombinations func(current string, start int)
	generateCombinations = func(current string, start int) {
		if len(current) == n {
			result = append(result, current)
			return
		}

		for i := start; i < len(digits); i++ {
			generateCombinations(current+string(digits[i]), i+1)
		}
	}

	generateCombinations("", 0)
	return result
}

// func main() {
// 	fmt.Println(MagicGrowthN(1))  // ["1", "2", "3", "4", "5", "6", "7", "8", "9"]
// 	fmt.Println(MagicGrowthN(2))  // ["01", "02", ... "08", "09", "12", "13" ... "78", "79", "89"]
// 	fmt.Println(MagicGrowthN(3))  // ["012", "013", ... "089", "123", "123" ... "678", "679", "789"]
// 	fmt.Println(MagicGrowthN(9))  // ["012345678", "012345679", "012345689", ..., "123456789"]
// 	fmt.Println(MagicGrowthN(10)) // ["0123456789"]
// }
