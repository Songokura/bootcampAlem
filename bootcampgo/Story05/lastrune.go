package bootcamp

// import "github.com/alem-platform/ap"

func LastRune(s string) rune {
	if s == "" {
		return rune(0)
	}
	runes := []rune(s)

	return runes[len(runes)-1]
}

// func main() {
// 	ap.PutRune(LastRune("salem"))
// 	ap.PutRune(LastRune("student"))
// 	ap.PutRune('\n')
// }
