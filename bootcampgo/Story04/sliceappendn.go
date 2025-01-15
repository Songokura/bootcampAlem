package bootcamp

func SliceAppendN(n int) []int {
	if n <= 0 {
		return []int{}
	}
	var slice []int
	for i := 0; i < n; i++ {
		slice = append(slice, i)
	}
	return slice
}
