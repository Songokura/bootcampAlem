package utils

// This function types the first line of the map
func MakeFirstLineOfBaseArray(arr *[][]rune) {
	for f := 0; f < len((*arr)[0])-1; f++ {
		if f == 0 || f == len((*arr)[0])-1 {
			(*arr)[0][f] = ' '
		} else {
			(*arr)[0][f] = '_'
		}
	}
}

// This function fills the default map in accordance with values of h and w
func MakeBaseArray(arr *[][]rune, h int) {
	for i := 1; i < h*3+1; i++ {
		for j := 0; j < len((*arr)[0]); j++ {
			(*arr)[i][j] = ' '
			if i%3 == 0 && j != 0 {
				(*arr)[i][j] = '_'
			}
			if j%8 == 0 || j == 0 {
				(*arr)[i][j] = '|'
			}
		}
	}
}

// This function puts entities on the map. It receives values of Walls, Player and Reward from input.
func FillEntity(arr *[][]rune, entities rune, x int, y int, type_character string) {
	x0 := 4 + y*8
	y0 := 2 + x*3
	if entities == '0' {
		for i := y0 - 1; i < y0+2; i++ {
			for j := x0 - 3; j < x0+4; j++ {
				(*arr)[i][j] = rune(type_character[0])
			}
		}
	} else if entities == '2' {
		(*arr)[y0][x0] = rune(type_character[1])
	} else if entities == '3' {
		(*arr)[y0][x0] = rune(type_character[2])
	} else {
		return
	}
}
