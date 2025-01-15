package bootcamp

func SliceMakeN(n int) []int {
	if n <= 0 {
		return []int{}
	}
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = i
	}
	return slice
}
