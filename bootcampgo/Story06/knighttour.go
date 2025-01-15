package bootcamp

var moves = [][2]int{
	{2, 1},
	{1, 2},
	{-1, 2},
	{-2, 1},
	{-2, -1},
	{-1, -2},
	{1, -2},
	{2, -1},
}

// KnightTour finds all solutions to the knight's tour problem for a given board size
func KnightTour(size int) [][][]int {
	if size <= 0 {
		return nil
	}

	// Initialize the board with zeros
	board := make([][]int, size)
	for i := range board {
		board[i] = make([]int, size)
	}

	var solutions [][][]int
	// Start the tour from the first position
	solveKnightTour(0, 0, 1, size, board, &solutions)

	return solutions
}

func solveKnightTour(x, y, moveNum, size int, board [][]int, solutions *[][][]int) {
	// Mark the current position with the move number
	board[x][y] = moveNum

	// If all positions are visited, add the solution
	if moveNum == size*size {
		*solutions = append(*solutions, deepCopy(board))
		// Unmark the current position and return
		board[x][y] = 0
		return
	}

	// Try all possible moves for a knight
	for _, move := range moves {
		nextX, nextY := x+move[0], y+move[1]
		if isValidMove(nextX, nextY, size, board) {
			solveKnightTour(nextX, nextY, moveNum+1, size, board, solutions)
		}
	}

	// Unmark the current position for backtracking
	board[x][y] = 0
}

func isValidMove(x, y, size int, board [][]int) bool {
	return x >= 0 && y >= 0 && x < size && y < size && board[x][y] == 0
}

func deepCopy(board [][]int) [][]int {
	size := len(board)
	copy := make([][]int, size)
	for i := range board {
		copy[i] = make([]int, size)
		for j := range board[i] {
			copy[i][j] = board[i][j]
		}
	}
	return copy
}
