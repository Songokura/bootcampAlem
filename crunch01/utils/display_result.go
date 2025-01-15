package utils

import "github.com/alem-platform/ap"

// This functions displays end looks of the map with colors and preferable characters for Walls, Player and Reward.
func DisplayRune(arr [][]rune) {
	Reset := []rune("\033[0m")
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if arr[i][j] == '|' {
				for n := j; n < len(arr[0]); n++ {
					if arr[i][n] == '|' {
						for _, val := range Reset {
							ap.PutRune(val)
						}
						ap.PutRune(arr[i][n])
					} else {
						SetColor(arr[i][n])
					}
					j++
				}
			} else {
				for _, val := range Reset {
					ap.PutRune(val)
				}
				ap.PutRune(arr[i][j])
			}
		}
		ap.PutRune('\n')
	}
}
