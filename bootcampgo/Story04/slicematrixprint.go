package bootcamp

import "github.com/alem-platform/ap"

func SliceMatrixPrint(matrix [][]rune) {
	numRows := len(matrix)
	for i, row := range matrix {
		numCols := len(row)
		for j, element := range row {
			ap.PutRune(element)
			if j < numCols-1 {
				ap.PutRune(' ')
			}
		}
		if i < numRows-1 {
			ap.PutRune('\n')
		}
	}
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
// }
