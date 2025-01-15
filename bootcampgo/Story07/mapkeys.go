package bootcamp

func MapKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// func main() {
// 	m := map[string]int{"one": 1, "two": 2, "three": 3}
// 	keys := MapKeys(m)
// 	fmt.Println(keys) // [one two three]
// }
