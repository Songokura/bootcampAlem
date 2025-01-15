package bootcamp

// import "fmt"

func solve_IslandCost(matrix *[][]int, x, y int, cost *int) {
	if (*matrix)[x][y] == 0 {
		return
	}

	*cost += (*matrix)[x][y]
	(*matrix)[x][y] = 0
	if x+1 < len(*matrix) {
		if (*matrix)[x+1][y] != 0 {
			solve_IslandCost(matrix, x+1, y, cost)
		}
	}
	if x-1 >= 0 {
		if (*matrix)[x-1][y] != 0 {
			solve_IslandCost(matrix, x-1, y, cost)
		}
	}
	if y+1 < len((*matrix)[x]) {
		if (*matrix)[x][y+1] != 0 {
			solve_IslandCost(matrix, x, y+1, cost)
		}
	}
	if y-1 >= 0 {
		if (*matrix)[x][y-1] != 0 {
			solve_IslandCost(matrix, x, y-1, cost)
		}
	}
}

func IslandCost(matrix [][]int, x, y int) int {
	if y >= len(matrix) || y < 0 {
		return 0
	}
	if x >= len(matrix[0]) || x < 0 {
		return 0
	}
	var sum int = 0
	solve_IslandCost(&matrix, y, x, &sum)
	return sum
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

// 	fmt.Println(IslandCost(matrix, 4, 2)) // 11
// 	fmt.Println(IslandCost(matrix, 0, 0)) // 5
// 	fmt.Println(IslandCost(matrix, 1, 4)) // 1
// 	fmt.Println(IslandCost(matrix, 0, 3)) // 0
// }
