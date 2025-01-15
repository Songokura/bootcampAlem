package main

import (
	"crunch01/enum"
	"crunch01/utils"
	"fmt"
)

func main() {
	var h int
	var w int
	var tmp string

	// Scanning height and width
	putHeightnWidth := enum.EnterHnW
	utils.Print(putHeightnWidth)
	fmt.Scanf("%d %d", &h, &w)
	for h <= 0 && w <= 0 {
		utils.Print(enum.ErrorHnW)
		utils.Print(putHeightnWidth)
		fmt.Scanf("%d %d", &h, &w)
	}

	// Creating array
	arr := make([][]rune, h*3+1)
	for i := range arr {
		arr[i] = make([]rune, (w*7)+w+1)
	}

	entities := make([]string, h)

	putStruct := enum.EnterStruct
	utils.Print(putStruct)
	// Scanning enteties
	for i := 0; i < h; i++ {
		fmt.Scanf("%s", &entities[i])
		if len(entities[i]) != w { // Checks for input errors
			utils.Print(enum.ErrorStruct)
			utils.Print(putStruct)
			i--
		} else if utils.IsValidStruct(entities[i]) == false {
			utils.Print(enum.ErrorNotInt)
			utils.Print(putStruct)
			i--
		}
	}

	var typeCharacter string = "X>@" // Default values

	putSymbols := enum.EnterSymbols
	utils.Print(putSymbols)

	// Scanning values to change enteties
	fmt.Scanf("%s", &tmp)
	for len(tmp) != 3 && len(tmp) != 0 {
		utils.Print(enum.ErrorChar)
		fmt.Scanf("%s", &tmp)
	}
	if len(tmp) == 3 {
		typeCharacter = tmp
	}

	// Fills the map with initial looks
	utils.MakeFirstLineOfBaseArray(&arr)
	utils.SetGlobalCharactersForColor(typeCharacter)
	utils.MakeBaseArray(&arr, h)

	// Filling array with enteties
	for i := 0; i < h; i++ {
		for j, v := range entities[i] {
			utils.FillEntity(&arr, v, i, j, typeCharacter)
		}
	}

	final_map_text := enum.FinalMap
	utils.Print(final_map_text)
	// Printing final Map
	utils.DisplayRune(utils.BuildNotatedMap(arr))
}
