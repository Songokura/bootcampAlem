package util

import "math/rand"

func RelocateBomb(row, col int, array [][]string, newArray [][]int) {
	array[row][col] = "."
	newArray[row][col] = 0

	// Find a random empty cell to place the bomb
	for {
		randRow := rand.Intn(len(array))
		randCol := rand.Intn(len(array[0]))
		if array[randRow][randCol] == "." {
			array[randRow][randCol] = "*"
			newArray[randRow][randCol] = 1
			break
		}
	}
}

func GenerateRandomMap(row, column int) ([][]string, [][]int, int) {
	array := make([][]string, row)
	newArray := make([][]int, row)
	bombs := 0

	// Generate random map
	PutString("Generating random map...\n")
	for i := 0; i < row; i++ {
		array[i] = make([]string, column)
		newArray[i] = make([]int, column)
		for j := 0; j < column; j++ {
			if rand.Intn(100) < 30 { // 30% chance of placing a bomb
				array[i][j] = "*"
				newArray[i][j] = 1 // Mine
				bombs++
			} else {
				array[i][j] = "."
				newArray[i][j] = 0 // Empty space
			}
		}
	}

	return array, newArray, bombs
}
