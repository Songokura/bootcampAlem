package bootcamp

func GameCities(cities []string) []string {
	used := make(map[string]bool)
	ans := []string{}
	temp := []string{}

	solve_GameCities(&ans, used, cities, temp)

	return ans
}

func solve_GameCities(ans *[]string, used map[string]bool, cities []string, temp []string) {
	if len(temp) > len(*ans) {
		*ans = make([]string, len(temp))
		copy(*ans, temp)
	}

	for _, city := range cities {
		last := getLast(temp)
		if !used[city] && (last == "" || toUpper(last[len(last)-1]) == toUpper(city[0])) {
			used[city] = true
			temp = append(temp, city)
			solve_GameCities(ans, used, cities, temp)
			temp = temp[:len(temp)-1]
			used[city] = false
		}
	}
}

func getLast(s []string) string {
	if len(s) == 0 {
		return ""
	}
	return s[len(s)-1]
}

func toUpper(s byte) byte {
	if s >= 'a' && s <= 'z' {
		return s - 32
	}
	return s
}

// func main() {
// 	cities := []string{
// 		"Astana",
// 		"Tokyo",
// 		"Amman",
// 		"Monaco",
// 		"Seoul",
// 		"Nassau",
// 		"Oslo",
// 		"Rabat",
// 		"Dover",
// 		"Ottawa",
// 		"Accra",
// 		"Ulaanbaatar",
// 	}
// 	chain := GameCities(cities)
// 	fmt.Println(chain) // [Monaco Oslo Ottawa Astana Accra Amman Nassau Ulaanbaatar Rabat Tokyo]
// }
