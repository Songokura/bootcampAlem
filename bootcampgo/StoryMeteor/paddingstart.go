package bootcamp

func PaddingStart(s string, totalLength int) string {
	curretLength := runeCount(s)

	paddingNeeded := totalLength - curretLength
	if paddingNeeded > 0 {
		padding := ""

		for i := 0; i < paddingNeeded; i++ {
			padding += " "
		}
		return padding + s
	}
	return s
}

func runeCount(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

// func main() {
// 	fmt.Printf("%q\n", PaddingStart("salem", 10))   // "     salem"
// 	fmt.Printf("%q\n", PaddingStart("salem 😀", 10)) // "   salem 😀"
// 	fmt.Printf("%q\n", PaddingStart("salem", 1))    // "salem"
// }
