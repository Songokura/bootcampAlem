package bootcamp

func RunesLen(arr [20]rune) int {
	lenght := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == '\x00' {
			break
		}
		lenght++
	}
	return lenght
}
