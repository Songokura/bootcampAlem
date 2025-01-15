package bootcamp

func MapGet(m map[string]int, key string) int {
	if val, ok := m[key]; ok {
		return val
	}
	return 0
}

// func main() {
// 	m := map[string]int{"one": 1, "two": 2, "three": 3}
// 	fmt.Println(MapGet(m, "two"))  // 2
// 	fmt.Println(MapGet(m, "four")) // 0
// }
