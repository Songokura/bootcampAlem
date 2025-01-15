package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func printInvalid() {
	ap.PutRune('N')
	ap.PutRune('V')
	ap.PutRune('\n')
}

func printInt(n int) {
	if n == 0 {
		ap.PutRune('0')
		return
	}
	digits := []rune{}
	for n > 0 {
		digits = append([]rune{rune('0' + n%10)}, digits...)
		n /= 10
	}
	for _, digit := range digits {
		ap.PutRune(digit)
	}
}

func getArg() (string, bool) {
	args := os.Args
	if len(args) != 2 {
		return "", false
	}
	return args[1], true
}

func atoi(s string) (int, bool) {
	n := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, false
		}
		n = n*10 + int(c-'0')
	}
	return n, true
}

func main() {
	arg, ok := getArg()
	if !ok {
		printInvalid()
		return
	}

	seconds, ok := atoi(arg)
	if !ok {
		printInvalid()
		return
	}

	const (
		secondsInMinute = 60
		secondsInHour   = 3600
		secondsInDay    = 86400
		maxInt32        = 2147483647
	)

	if seconds > maxInt32 {
		printInvalid()
		return
	}

	days := seconds / secondsInDay
	seconds %= secondsInDay

	hours := seconds / secondsInHour
	seconds %= secondsInHour

	minutes := seconds / secondsInMinute
	seconds %= secondsInMinute

	printed := false
	if days > 0 {
		printInt(days)
		ap.PutRune('d')
		printed = true
	}
	if hours > 0 {
		if printed {
			ap.PutRune(' ')
		}
		printInt(hours)
		ap.PutRune('h')
		printed = true
	}
	if minutes > 0 {
		if printed {
			ap.PutRune(' ')
		}
		printInt(minutes)
		ap.PutRune('m')
		printed = true
	}
	ap.PutRune(' ')
	if seconds > 0 || !printed {
		printInt(seconds)
		ap.PutRune('s')
	}
	ap.PutRune('\n')
}
