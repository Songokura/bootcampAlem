package bootcamp

func CountWords(s string) map[string]int {
	Count := make(map[string]int)

	currWord := ""

	for _, char := range s {
		lowerChar := rune(char | ' ')

		if lowerChar >= 'a' && lowerChar <= 'z' {
			currWord += string(lowerChar)
		} else if currWord != "" {
			Count[currWord]++
			currWord = ""
		}
	}
	if currWord != "" {
		Count[currWord]++
	}
	return Count
}

// func main() {
// 	s := "The soup was stirred and stirred until thickened."
// 	wordCounts := CountWords(s)
// 	fmt.Println(wordCounts) // map[the:1 soup:1 was:1 and:1 stirred:2 until:1 thickened:1]
// }
