package bootcamp

// import "fmt"

func SliceUnion(slices ...[]int) []int {
	var doc []int
	var add bool
	for _, slice := range slices {
		for _, v := range slice {
			add = true

			for _, val_res := range doc {
				if val_res == v {
					add = false
					break
				}
			}
			if add {
				doc = append(doc, v)
			}
		}
	}
	return doc
}

// func main() {
// 	result := SliceUnion([]int{1, 2, 3, 20}, []int{1, 2, 3, 4}, []int{15, 0, 2})
// 	fmt.Println(result) // [1, 2, 3, 20, 4, 5, 15, 0]
// }
