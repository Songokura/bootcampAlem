package utils

/*
writen by tkoszhan
*/

func BuildNotatedMap(arr [][]rune) [][]rune {
	// Creating bigger array
	notatedMap := make([][]rune, len(arr)+1)

	extraColumnCount := lenOfNum(len(arr)) - 3 // How many extra column we should add
	if extraColumnCount < 0 {
		extraColumnCount = 0
	}
	for i := range notatedMap {
		notatedMap[i] = make([]rune, len(arr[0])+3+extraColumnCount)
	}

	// Filling letter

	pos := 7 + extraColumnCount
	notatedLetters := []rune{'@'} // Ascii symbol before 'A'
	for i := 0; i < len(notatedMap[0]); i++ {
		if i == pos {
			notatedLetters = nextLetter(notatedLetters)
			//
			for q, v := range notatedLetters {
				center := len(notatedLetters) / 2
				notatedMap[0][i+q-center] = v
			}
			pos += 8
		}
		if notatedMap[0][i] == 0 {
			notatedMap[0][i] = ' '
		}
	}

	// Filling numbers
	pos = 3
	numbers := []rune{'0'}

	for i := 1; i < len(notatedMap); i++ {
		for j := 0; j < 3+extraColumnCount; j++ {
			if i == pos && j == 0 {
				numbers = nextNumber(numbers)
				for q, v := range numbers {
					notatedMap[i][j+q] = v
				}
				pos += 3
			}
			if notatedMap[i][j] == 0 {
				notatedMap[i][j] = ' '
			}
		}
	}

	// filling new array with old array values
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			notatedMap[i+1][j+3+extraColumnCount] = arr[i][j]
		}
	}

	return notatedMap
}

func nextLetter(notatedLetters []rune) []rune {
	n := len(notatedLetters)

	// Starting from end of array and incresing letter if element is zero
	for i := n - 1; i >= 0; i-- {
		if notatedLetters[i] == 'Z' {
			notatedLetters[i] = 'A'
		} else {
			notatedLetters[i]++   // Like 'B' -> 'C'
			return notatedLetters // Return after moved to next letter and it wasn't Z 
		}
	}

	// If all positions was Z we add A to the begining
	notatedLetters = append([]rune{'A'}, notatedLetters...)

	return notatedLetters
}

// Just like NextColor
func nextNumber(notatedNumbers []rune) []rune {
	n := len(notatedNumbers)

	for i := n - 1; i >= 0; i-- {
		if notatedNumbers[i] == '9' {
			notatedNumbers[i] = '0'
		} else {
			notatedNumbers[i]++
			return notatedNumbers
		}
	}

	notatedNumbers = append([]rune{'1'}, notatedNumbers...)

	return notatedNumbers
}

// Just returns counts len of int
func lenOfNum(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}
