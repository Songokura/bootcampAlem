package bootcamp

func ArraySetOne(arr *[20]int, idx int, val int) bool {
	if idx < 0 || idx >= len(arr) {
		return false
	}
	(*arr)[idx] = val
	return true
}
