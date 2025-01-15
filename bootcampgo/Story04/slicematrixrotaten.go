package bootcamp

import "github.com/alem-platform/ap"

// func SliceMatrixPrint(matrix [][]rune) {
// 	numRows := len(matrix)
// 	for i, row := range matrix {
// 		numCols := len(row)
// 		for j, element := range row {
// 			ap.PutRune(element)
// 			if j < numCols-1 {
// 				ap.PutRune(' ')
// 			}
// 		}
// 		if i < numRows-1 {
// 			ap.PutRune('\n')
// 		}
// 	}
// }

func SliceMatrixRotateN(matrix [][]rune, turns int) {
	if !isSquareMatrix(matrix) {
		return
	}

	n := len(matrix)
	turns = (turns%4 + 4) % 4 // Ensure turns is between 0 and 3

	for t := 0; t < turns; t++ {
		for i := 0; i < n/2; i++ {
			for j := i; j < n-i-1; j++ {
				temp := matrix[i][j]
				matrix[i][j] = matrix[n-1-j][i]
				matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
				matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
				matrix[j][n-1-i] = temp
			}
		}
	}
	ap.PutRune('\n')
}

func isSquareMatrix(matrix [][]rune) bool {
	rows := len(matrix)
	if rows == 0 {
		return false
	}
	cols := len(matrix[0])
	return rows == cols
}

// func main() {
// 	matrix := [][]rune{
// 		{'a', 'b', 'c'},
// 		{'d', 'e', 'f'},
// 		{'g', 'h', 'i'},
// 	}

// 	SliceMatrixPrint(matrix)
// 	// Output:
// 	// a b c
// 	// d e f
// 	// g h i

// 	SliceMatrixRotateN(matrix, 1)

// 	SliceMatrixPrint(matrix)
// 	// Output:
// 	// g d a
// 	// h e b
// 	// i f c

// 	SliceMatrixRotateN(matrix, -2)

// 	SliceMatrixPrint(matrix)
// 	// Output:
// 	// c f i
// 	// b e h
// 	// a d g
// }
