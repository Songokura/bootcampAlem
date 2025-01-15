package bootcamp

func RepeatEcho(s string) string {
	var result string
	i := 0
	for i < len(s) {
		if isDigit(s[i]) {
			numStart := i
			for i < len(s) && isDigit(s[i]) {
				i++
			}
			num := convertToInt(s[numStart:i])

			if i < len(s) && s[i] == '(' {
				openCount := 1
				j := i + 1
				for j < len(s) && openCount > 0 {
					if s[j] == '(' {
						openCount++
					} else if s[j] == ')' {
						openCount--
					}
					j++
				}
				if openCount == 0 {
					pattern := s[i+1 : j-1]
					result += repeatPattern(pattern, num)
					i = j
					continue
				}
			}
		}
		result += string(s[i])
		i++
	}
	return result
}

func repeatPattern(pattern string, times int) string {
	var result string
	for i := 0; i < times; i++ {
		result += RepeatEcho(pattern)
	}
	return result
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func convertToInt(s string) int {
	num := 0
	for _, ch := range s {
		num = num*10 + int(ch-'0')
	}
	return num
}

// // Main function to test the RepeatEcho function
// func main() {
// 	fmt.Println(RepeatEcho("Ba2(na), 2(d2(a)) 10(!)(")) // Output: Banana daadaa !!!!!!!!!!(
// }
