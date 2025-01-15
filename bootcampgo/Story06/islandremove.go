package bootcamp

// func PrintMatrix(matrix [][]int) {
// 	for _, row := range matrix {
// 		for _, val := range row {
// 			fmt.Printf("%d ", val)
// 		}
// 		fmt.Println()
// 	}
// 	fmt.Println()
// }

var islandOffset [][]int = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func IslandRemove(matrix [][]int, x, y int) {
	height := MatrixHeight(matrix)
	width := MatrixWidth(matrix)
	IslandRemoveWithSize(matrix, width, height, x, y)
}

func IslandRemoveWithSize(matrix [][]int, width, height, x, y int) {
	if ExceedsBorder(matrix, width, height, x, y) {
		return
	}
	if matrix[y][x] == 0 {
		return
	}
	matrix[y][x] = 0

	for _, offset := range islandOffset {
		xOffset := offset[0]
		yOffset := offset[1]
		IslandRemoveWithSize(matrix, width, height, x+xOffset, y+yOffset)
	}
}

func ExceedsBorder(matrix [][]int, width, height, x, y int) bool {
	if x < 0 || y < 0 {
		return true
	}
	return x >= width || y >= height
}

func MatrixHeight(matrix [][]int) int {
	return len(matrix)
}

func MatrixWidth(matrix [][]int) int {
	widthFound := false
	width := -1
	for _, subArray := range matrix {
		if !widthFound {
			width = len(subArray)
			widthFound = true
			continue
		}
		if width != len(subArray) {
			return width
		}
	}
	return width
}

// func main() {
// 	matrix := [][]int{
// 		{1, 1, 1, 0, 0, 0, 0, 0, 0},
// 		{1, 1, 0, 0, 1, 0, 0, 0, 0},
// 		{0, 0, 0, 1, 2, 3, 1, 0, 0},
// 		{0, 0, 0, 1, 1, 1, 0, 0, 1},
// 		{0, 1, 0, 0, 0, 0, 0, 1, 2},
// 		{0, 0, 0, 1, 0, 0, 0, 0, 1},
// 		{0, 0, 1, 1, 2, 2, 0, 0, 0},
// 	}
// 	PrintMatrix(matrix)
// 	// Output:
// 	// 1 1 1 0 0 0 0 0 0
// 	// 1 1 0 0 1 0 0 0 0
// 	// 0 0 0 1 2 3 1 0 0
// 	// 0 0 0 1 1 1 0 0 1
// 	// 0 1 0 0 0 0 0 1 2
// 	// 0 0 0 1 0 0 0 0 1
// 	// 0 0 1 1 2 2 0 0 0

// 	IslandRemove(matrix, 4, 2)

// 	PrintMatrix(matrix)
// 	// Output:
// 	// 1 1 1 0 0 0 0 0 0
// 	// 1 1 0 0 0 0 0 0 0
// 	// 0 0 0 0 0 0 0 0 0
// 	// 0 0 0 0 0 0 0 0 1
// 	// 0 1 0 0 0 0 0 1 2
// 	// 0 0 0 1 0 0 0 0 1
// 	// 0 0 1 1 2 2 0 0 0
// }
