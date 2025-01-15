package utils

import (
	"github.com/alem-platform/ap"
)

var (
	Wall   = 'X'
	Player = '>'
	Reward = '@'
)

// Reveives rune and prints it with the color
func SetColor(v rune) {
	bgRed := []rune("\033[41m")
	bgGreen := []rune("\033[42m")
	bgYellow := []rune("\033[43m")
	bgWhite := []rune("\033[47m")
	switch v {
	case Wall:
		bgRed = append(bgRed, Wall)

		for _, val := range bgRed {
			ap.PutRune(val)
		}
	case Player:
		bgGreen = append(bgGreen, Player)
		for _, val := range bgGreen {
			ap.PutRune(val)
		}
	case Reward:
		bgYellow = append(bgYellow, Reward)

		for _, val := range bgYellow {
			ap.PutRune(val)
		}
	case ' ', '_':
		bgWhite = append(bgWhite, ' ')
		for _, val := range bgWhite {
			ap.PutRune(val)
		}
	default:
		ap.PutRune(v)
	}
}

// Sets values for wall, Player and Reward if they have been changed
func SetGlobalCharactersForColor(typeCharacter string) {
	for i, val := range typeCharacter {
		switch i {
		case 0:
			Wall = val
		case 1:
			Player = val
		case 2:
			Reward = val
		}
	}
}
