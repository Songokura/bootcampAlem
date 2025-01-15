package bootcamp

func ArraySlice(arr [20]int, low int, high int) []int {
	if low < 0 || low >= len(arr) || high < 0 || high > len(arr) || low > high {
		return []int{}
	}
	return arr[low:high]
}
