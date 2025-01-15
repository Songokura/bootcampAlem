package util

import "github.com/alem-platform/ap"

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
)

func PrintMap(row, column int, array [][]string, gameover bool) {
	PrintSpaces(5, ' ')
	// Print the column numbers
	for i := 0; i < column; i++ {
		if i == 0 {
			PrintSpaces(4, ' ')
			ap.PutRune(rune(i+1) + '0')
		} else {
			PrintSpaces(7, ' ')
			ap.PutRune(rune(i+1) + '0')
		}
	}
	PutRune('\n')
	// Print top border
	PutRune(' ')
	PrintSpaces(5, ' ')
	HeaderPrint(column)
	PutRune('\n')
	for i := 0; i < row; i++ {
		for j := 0; j < 3; j++ {
			// Print the row numbers
			switch j {
			case 0:
				PrintSpaces(5, ' ')
			case 1:
				PrintSpaces(2, ' ')
				PutRune(rune(i+1) + '0')
				PrintSpaces(2, ' ')
			case 2:
				PrintSpaces(5, ' ')
			}

			for k := 0; k < column; k++ {
				switch array[i][k] {
				case "*":
					if gameover {
						PrintCenter(j, '*', White, Reset)
					} else {
						PrintWall('X')
					}
				case "0":
					PrintCenter(j, ' ', White, Reset)
				case "1":
					PrintCenter(j, '1', Blue, Reset)
				case "2":
					PrintCenter(j, '2', Green, Reset)
				case "3":
					PrintCenter(j, '3', Red, Reset)
				case "4":
					PrintCenter(j, '4', Yellow, Reset)
				case "5":
					PrintCenter(j, '5', Magenta, Reset)
				case "6":
					PrintCenter(j, '6', Cyan, Reset)
				case "7":
					PrintCenter(j, '7', Blue, Reset)
				case "8":
					PrintCenter(j, '8', Green, Reset)
				default:
					PrintWall('X')
				}
			}
			PutRune('|')
			PutRune('\n')
		}
	}
}

// First row of the block is consists of the spaces, second row has the center element
func PrintCenter(column int, r rune, color, reset string) {
	if column == 0 {
		PutRune('|')
		PrintSpaces(7, ' ')
	} else if column == 1 {
		PutRune('|')
		PrintSpaces(3, ' ')
		PutString(color)
		ap.PutRune(r)
		PutString(reset)
		PrintSpaces(3, ' ')
	} else {
		PutRune('|')
		PrintSpaces(7, '_')
	}
}
