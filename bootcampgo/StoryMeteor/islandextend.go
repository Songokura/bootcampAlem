package bootcamp

// import "fmt"

func IslandExtend(matrix [][]int, x, y int) bool {
	if x < 0 || y < 0 || x >= len(matrix[0]) || y >= len(matrix) || matrix[y][x] == 0 {
		return false
	}
	// arr := []bool{IslandExtend(matrix, x-1, y), IslandExtend(matrix, x+1, y), IslandExtend(matrix, x, y-1), IslandExtend(matrix, x, y+1)}
	// for _, v := range arr {
	//  if v == true {
	//   return true
	//  }
	// }

	bool_matrix := make([][]bool, len(matrix))
	for i := range bool_matrix {
		bool_matrix[i] = make([]bool, len(matrix[0]))
		for j := range bool_matrix[i] {
			bool_matrix[i][j] = false // Initialize each element to false
		}
	}
	// fmt.Println(bool_matrix)
	return solveIsland(matrix, bool_matrix, x, y)
}

func solveIsland(matrix [][]int, bool_matrix [][]bool, x, y int) bool {
	if x < 0 || y < 0 || x >= len(matrix[0]) || y >= len(matrix) || bool_matrix[y][x] == true {
		return false
	}
	if matrix[y][x] == 0 {
		matrix[y][x] = 1
		bool_matrix[y][x] = true
		return true
	}

	bool_matrix[y][x] = true

	arr := []bool{solveIsland(matrix, bool_matrix, x-1, y), solveIsland(matrix, bool_matrix, x+1, y), solveIsland(matrix, bool_matrix, x, y-1), solveIsland(matrix, bool_matrix, x, y+1)}
	for _, v := range arr {
		if v == true {
			return true
		}
	}
	return false
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

// 	fmt.Println(IslandExtend(matrix, 4, 2)) // true
// 	print(matrix)
// 	// 1 1 1 0 1 0 0 0 0
// 	// 1 1 0 1 1 1 1 0 0
// 	// 0 0 1 1 2 3 1 1 0
// 	// 0 0 1 1 1 1 1 0 1
// 	// 0 1 0 1 1 1 0 1 2
// 	// 0 0 0 1 0 0 0 0 1
// 	// 0 0 1 1 2 2 0 0 0

// 	fmt.Println(IslandExtend(matrix, 0, 0)) // true
// 	print(matrix)
// 	// 1 1 1 1 1 0 0 0 0
// 	// 1 1 1 1 1 1 1 0 0
// 	// 1 1 1 1 2 3 1 1 0
// 	// 0 0 1 1 1 1 1 0 1
// 	// 0 1 0 1 1 1 0 1 2
// 	// 0 0 0 1 0 0 0 0 1
// 	// 0 0 1 1 2 2 0 0 0

// 	fmt.Println(IslandExtend(matrix, 1, 4)) // true
// 	print(matrix)
// 	// 1 1 1 1 1 0 0 0 0
// 	// 1 1 1 1 1 1 1 0 0
// 	// 1 1 1 1 2 3 1 1 0
// 	// 0 1 1 1 1 1 1 0 1
// 	// 1 1 1 1 1 1 0 1 2
// 	// 0 1 0 1 0 0 0 0 1
// 	// 0 0 1 1 2 2 0 0 0

// 	fmt.Println(IslandExtend(matrix, 0, 3)) // false
// 	print(matrix)
// 	// 1 1 1 1 1 0 0 0 0
// 	// 1 1 1 1 1 1 1 0 0
// 	// 1 1 1 1 2 3 1 1 0
// 	// 0 1 1 1 1 1 1 0 1
// 	// 1 1 1 1 1 1 0 1 2
// 	// 0 1 0 1 0 0 0 0 1
// 	// 0 0 1 1 2 2 0 0 0
// }

// func print(arr [][]int) {
// 	for _, v := range arr {
// 		for _, j := range v {
// 			fmt.Print(j)
// 			fmt.Print(" ")
// 		}
// 		fmt.Println()
// 	}
// }
