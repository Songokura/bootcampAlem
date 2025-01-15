package bootcamp

func NsthRune(s string, n int) rune {
	if s == "" {
		return rune(0)
	}
	runes := []rune(s)

	return runes[n]
}

// func main() {
// 	ap.PutRune(NthRune("salem", 0))
// 	ap.PutRune(NthRune("student", 2))
// 	ap.PutRune('\n')
// }
