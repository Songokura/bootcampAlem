package bootcamp

func MapEqual(m1 map[string]int, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for key, val1 := range m1 {
		val2, ok := m2[key]
		if !ok || val1 != val2 {
			return false
		}
	}

	return true
}

// func main() {
// 	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
// 	m2 := map[string]int{"one": 1, "two": 2, "three": 3}
// 	m3 := map[string]int{"one": 1, "two": 2, "four": 4}

// 	fmt.Println(MapEqual(m1, m2)) // true
// 	fmt.Println(MapEqual(m1, m3)) // false
// }
