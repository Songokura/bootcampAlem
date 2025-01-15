package util

func ScanBombs(c1, c2 int, array [][]string, dots int, gameover bool, win bool) (int, int, bool, bool) {
	if c1 < 0 || c2 < 0 || c1 >= len(array) || c2 >= len(array[0]) {
		return 0, dots, gameover, win
	}

	// If the cell is already revealed or has a number, return 0
	if array[c1][c2] != "." {

		PutString("This cell is already open.\n")
		return Atoi(array[c1][c2]), dots, gameover, win
	}

	// Mark the cell as revealed
	array[c1][c2] = Itoa(CountAdjacentBombs(c1, c2, array))

	// If there are no adjacent bombs, recursively reveal neighboring cells
	if array[c1][c2] == "0" {
		for i := c1 - 1; i <= c1+1; i++ {
			for j := c2 - 1; j <= c2+1; j++ {
				ScanBombs(i, j, array, dots, gameover, win)
			}
		}
	}
	// count openCells
	openCells := 0
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			if array[i][j] != "." && array[i][j] != "*" {
				openCells++
			}
		}
	}

	if openCells >= dots {
		win = true
		gameover = true
	}

	// Return the count of adjacent bombs
	return Atoi(array[c1][c2]), dots, gameover, win
}

func CountAdjacentBombs(c1, c2 int, array [][]string) int {
	count := 0
	for i := c1 - 1; i <= c1+1; i++ {
		for j := c2 - 1; j <= c2+1; j++ {
			if i >= 0 && j >= 0 && i < len(array) && j < len(array[0]) && array[i][j] == "*" {
				count++
			}
		}
	}
	return count
}
