package bootcamp

// import "fmt"

func MagicGrowth2() []string {
	var combinations []string
	digits := "0123456789"
	for i := 0; i < 10; i++ {
		for j := i + 1; j < 10; j++ {
			combinations = append(combinations, string(digits[i])+string(digits[j]))
		}
	}
	return combinations
}

// func main() {
// 	fmt.Println(MagicGrowth2()) // ["01", "02", ... "08", "09", "12", "13" ... "78", "79", "89"]
// }
