package bootcamp

func EightQueens() [][]int {
	var solutions [][]int
	board := make([]int, 8)
	findSolutions(&solutions, board, 0)
	return solutions
}

func findSolutions(solutions *[][]int, board []int, row int) {
	if row == 8 {
		solution := make([]int, 8)
		copy(solution, board)
		*solutions = append(*solutions, solution)
		return
	}

	for col := 1; col < 9; col++ {
		if isSafe(board, row, col) {
			board[row] = col
			findSolutions(solutions, board, row+1)
			board[row] = 0
		}
	}
}

func isSafe(board []int, row, col int) bool {
	for r := 0; r < row; r++ {
		if board[r] == col {
			return false
		}
		if Abs(board[r]-col) == (row - r) {
			return false
		}
	}
	return true
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// func main() {
// 	solutions := EightQueens()
// 	for _, solution := range solutions {
// 		for _, pos := range solution {
// 			fmt.Print(pos)
// 		}
// 		fmt.Println()
// 	}
// }
