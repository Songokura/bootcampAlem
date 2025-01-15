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

func SliceMatrixRotateClockwise(matrix [][]rune) {
	n := len(matrix)

	// Check if the matrix is square
	if n != len(matrix[0]) {
		// Matrix is not square, return without rotating
		return
	}

	// Perform rotation layer by layer, from outer to inner layers
	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - 1 - layer
		for i := first; i < last; i++ {
			// Store top element
			top := matrix[first][i]

			// Move left element to top
			matrix[first][i] = matrix[n-1-i][first]

			// Move bottom element to left
			matrix[n-1-i][first] = matrix[last][n-1-i]

			// Move right element to bottom
			matrix[last][n-1-i] = matrix[i][last]

			// Move top element to right
			matrix[i][last] = top
		}
	}
	ap.PutRune('\n')
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

// 	SliceMatrixRotateClockwise(matrix)

// 	SliceMatrixPrint(matrix)
// 	// Output:
// 	// g d a
// 	// h e b
// 	// i f c
// }
