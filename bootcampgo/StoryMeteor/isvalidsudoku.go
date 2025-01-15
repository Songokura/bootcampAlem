package bootcamp

func ValidSudoku(n [9][9]int) bool {
	// Check each row
	for i := 0; i < 9; i++ {
		if !isValidSet(n[i][:]) {
			return false
		}
	}

	// Check each column
	for j := 0; j < 9; j++ {
		column := make([]int, 9)
		for i := 0; i < 9; i++ {
			column[i] = n[i][j]
		}
		if !isValidSet(column) {
			return false
		}
	}

	// Check each 3x3 subgrid
	for blockRow := 0; blockRow < 3; blockRow++ {
		for blockCol := 0; blockCol < 3; blockCol++ {
			block := make([]int, 0, 9)
			for i := blockRow * 3; i < blockRow*3+3; i++ {
				for j := blockCol * 3; j < blockCol*3+3; j++ {
					block = append(block, n[i][j])
				}
			}
			if !isValidSet(block) {
				return false
			}
		}
	}

	return true
}

// Helper function to check if a set of 9 numbers is valid (contains each number from 1 to 9 exactly once)
func isValidSet(set []int) bool {
	seen := make(map[int]bool)
	for _, num := range set {
		if num < 1 || num > 9 || seen[num] {
			return false
		}
		seen[num] = true
	}
	return true
}

// func main() {
// 	validSudoku := [9][9]int{
// 		{5, 3, 4, 6, 7, 8, 9, 1, 2},
// 		{6, 7, 2, 1, 9, 5, 3, 4, 8},
// 		{1, 9, 8, 3, 4, 2, 5, 6, 7},
// 		{8, 5, 9, 7, 6, 1, 4, 2, 3},
// 		{4, 2, 6, 8, 5, 3, 7, 9, 1},
// 		{7, 1, 3, 9, 2, 4, 8, 5, 6},
// 		{9, 6, 1, 5, 3, 7, 2, 8, 4},
// 		{2, 8, 7, 4, 1, 9, 6, 3, 5},
// 		{3, 4, 5, 2, 8, 6, 1, 7, 9},
// 	}

// 	invalidSudoku := [9][9]int{
// 		{5, 3, 4, 6, 7, 8, 9, 1, 2},
// 		{6, 7, 2, 1, 9, 5, 3, 4, 8},
// 		{1, 9, 8, 3, 4, 2, 5, 6, 7},
// 		{8, 5, 9, 7, 6, 1, 4, 2, 3},
// 		{4, 2, 6, 8, 5, 3, 7, 9, 1},
// 		{7, 1, 3, 9, 2, 4, 8, 5, 6},
// 		{9, 6, 1, 5, 3, 7, 2, 8, 4},
// 		{2, 8, 7, 4, 1, 9, 6, 3, 5},
// 		{3, 4, 5, 2, 8, 6, 1, 7, 8}, // invalid row
// 	}

// 	fmt.Println(ValidSudoku(validSudoku))   // true
// 	fmt.Println(ValidSudoku(invalidSudoku)) // false
// }
