package util

func PrintStatistics(row, column, bombs, moves int) {
	// If takes the row and column nums to represent rxc field size, prints the number of bombs, moves
	PutString("Your statistics: \n")
	PutString(" - Field size: ")
	PutString(Itoa(row))
	PutString("x")
	PutString(Itoa(column))
	PutString("\n - Number of bombs: ")
	PutString(Itoa(bombs))
	PutString("\n - Number of moves: ")
	PutString(Itoa(moves))
	PutRune('\n')
}
