package main

import (
	"crunch02/text"
	"crunch02/util"
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(42) // Fixed seed for deterministic randomness

	gameover := false
	moves := 0
	win := false
	var row, column int
	util.PutString(text.EnterInput)

	fmt.Scanf("%d %d\n", &row, &column)

	for row <= 2 || column <= 2 {
		util.PutString(text.InputError1)
		fmt.Scanf("%d %d\n", &row, &column)
	}

	var array [][]string
	var newArray [][]int
	var bombs int

	// Menu for map generation
	util.PutString(text.GridOp)
	util.PutString(text.GridOp1)
	util.PutString(text.GridOp2)

	var choice int
	fmt.Scanf("%d", &choice)

	switch choice {
	case 1:
		// Create a 2D slice to represent the map
		array = make([][]string, row)
		for i := range array {
			array[i] = make([]string, column)
		}
		newArray = make([][]int, row)
		for i := range newArray {
			newArray[i] = make([]int, column)
		}

		// Read the map
		util.PutString(text.StructGrid)
		bombs = 0
		for i := 0; i < row; {
			var line string
			fmt.Scanf("%s\n", &line)
			validLine := true
			if len(line) != column { // Check the lenght of the line, it must be equal to the column
				util.PutString(text.InputError2)
				validLine = false
				continue
			}
			for j, ch := range line {
				if ch == '.' {
					array[i][j] = string(ch)
					newArray[i][j] = 0 // Empty space
				} else if ch == '*' {
					array[i][j] = string(ch)
					newArray[i][j] = 1 // Mine
					bombs++
				} else {
					util.PutString(text.InputError3) // Input can only be '.' or '*'.
					validLine = false
					break
				}
			}
			if validLine {
				i++ // Increment i only if the input line is valid
			}
		}

		if bombs < 2 {
			util.PutString(text.InputError4) // Amount of bombs cannot be less than 2.
			gameover = true
			break
		}

	case 2:
		// Generate a random map
		array, newArray, bombs = util.GenerateRandomMap(row, column)
	default:
		util.PutString(text.InputError5)
		return
	}

	util.PutString(text.Map) // Your map:
	util.PrintMap(row, column, array, gameover)

	initialMove := true

	for !gameover {
		var coor1, coor2 int
		// Entering coordinates until they are valid
		for {
			fmt.Print(text.Coord)
			fmt.Scanf("%d %d", &coor1, &coor2)
			if coor1 < 1 || coor1 > len(array) || coor2 < 1 || coor2 > len(array[0]) {
				fmt.Println(text.InputError6)
				continue
			}
			moves++
			break
		}
		// dots = area of the map - amount of bombs
		dots := (row * column) - bombs

		if initialMove && array[coor1-1][coor2-1] == "*" {
			// Relocate bomb from the initial move
			util.RelocateBomb(coor1-1, coor2-1, array, newArray)
		}

		if array[coor1-1][coor2-1] == "*" { // If player hits the bomb the loop will be stopped
			gameover = true
			util.PrintMap(row, column, array, gameover)
			util.PutString(text.Lose)
			util.PrintStatistics(row, column, bombs, moves)
			break
		} else {
			var count int
			// If player chooses the coordinate that is not a bomb we calculate the amount of nearest bombs
			count, dots, gameover, win = util.ScanBombs(coor1-1, coor2-1, array, dots, gameover, win)
			array[coor1-1][coor2-1] = util.Itoa(count)
			// Printing of the map
			util.PrintMap(row, column, array, gameover)
			if win == true {
				gameover = true
				util.PutString(text.Win) // Congrats! You won!
				util.PrintStatistics(row, column, bombs, moves)
				break
			}
		}
		initialMove = false
	}
}
