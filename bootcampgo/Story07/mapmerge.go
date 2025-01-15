package bootcamp

func MapMerge(m1 map[string]int, m2 map[string]int) map[string]int {
	merge := make(map[string]int)

	for key, value := range m1 {
		merge[key] = value
	}
	for key, value := range m2 {
		merge[key] = value
	}
	return merge
}

// func main() {
// 	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
// 	m2 := map[string]int{"two": 22, "four": 4}
// 	merged := MapMerge(m1, m2)
// 	fmt.Println(merged) // map[one:1 two:22 three:3 four:4]
// }
