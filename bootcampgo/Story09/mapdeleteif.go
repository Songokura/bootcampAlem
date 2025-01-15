package bootcamp

func MapDeleteIf(m map[string]int, fn func(int) bool) {
	Delete := make([]string, 0)

	for key, value := range m {
		if fn(value) {
			Delete = append(Delete, key)
		}
	}

	for _, v := range Delete {
		delete(m, v)
	}
}

// func main() {
// 	isEven := func(n int) bool {
// 		return n%2 == 0
// 	}

// 	m := map[string]int{
// 		"a": 1,
// 		"b": 2,
// 		"c": 3,
// 		"d": 4,
// 	}

// 	MapDeleteIf(m, isEven)
// 	fmt.Println(m) // map[a:1 c:3]
// }
