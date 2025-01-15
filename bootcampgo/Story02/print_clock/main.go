package main

import "github.com/alem-platform/ap"

func PrintTime(hour, minute int) {
	ap.PutRune('0' + rune(hour/10))
	ap.PutRune('0' + rune(hour%10))

	ap.PutRune(':')

	ap.PutRune('0' + rune(minute/10))
	ap.PutRune('0' + rune(minute%10))

	ap.PutRune('\n')
}

func main() {
	for hour := 0; hour < 24; hour++ {
		for minute := 0; minute < 60; minute++ {
			PrintTime(hour, minute)
		}
	}
}
