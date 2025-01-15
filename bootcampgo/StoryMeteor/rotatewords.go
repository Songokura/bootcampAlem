package bootcamp

func RotateWords(s string) string {
	words := splitWords(s)
	if len(words) <= 1 {
		return s
	}
	lastWord := words[len(words)-1]
	rotatedWords := append([]string{lastWord}, words[:len(words)-1]...)
	return joinWords(rotatedWords)
}

func splitWords(s string) []string {
	var words []string
	var wordStart int

	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == ' ' {
			if i > wordStart {
				words = append(words, s[wordStart:i])
			}
			wordStart = i + 1
		}
	}

	return words
}

func joinWords(words []string) string {
	result := ""
	for i, word := range words {
		if i > 0 {
			result += " "
		}
		result += word
	}
	return result
}
