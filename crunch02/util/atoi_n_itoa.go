package util

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}
	result := ""
	for n > 0 {
		digit := n % 10
		result = string('0'+digit) + result
		n /= 10
	}
	if isNegative {
		result = "-" + result
	}
	return result
}

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	start := 0
	result := 0
	sign := 1
	if s[0] == '-' || s[0] == '+' {
		start++
		if s[0] == '-' {
			sign = -1
		}
	}
	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		result = result*10 + int(s[i]-'0')
	}
	return result * sign
}
