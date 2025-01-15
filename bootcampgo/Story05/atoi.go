package bootcamp

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	// Skip leading whitespace
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}

	// Check if the number is negative
	sign := 1
	if i < len(s) && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	// Convert string to integer
	num := 0
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0 // Invalid character found
		}
		digit := int(s[i] - '0')
		// Check for overflow
		if num > (1<<31-1)/10 || (num == (1<<31-1)/10 && digit > 7) {
			// Return max or min int32 value depending on sign
			if sign == 1 {
				return 1<<31 - 1
			} else {
				return -(1 << 31)
			}
		}
		num = num*10 + digit
	}

	return num * sign
}

// func trimSpace(s string) string {
// 	start, end := 0, len(s)-1

// 	// Trim leading whitespace
// 	for start <= end && s[start] == ' ' {
// 		start++
// 	}

// 	// Trim trailing whitespace
// 	for end >= start && s[end] == ' ' {
// 		end--
// 	}

// 	return s[start : end+1]
// }

// func main() {
// 	fmt.Println(Atoi("123"))
// 	fmt.Println(Atoi("+123"))
// 	fmt.Println(Atoi("-123$$$"))
// 	fmt.Println(Atoi("-123!"))
// 	fmt.Println(Atoi("abc"))
// }
