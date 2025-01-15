package utils

// This function checks correct input for entities, which is a map structure
func IsValidStruct(entities string) bool {
	for _, v := range entities {
		if v > '3' || v < '0' {
			return false
		}
	}
	return true
}

// Error return for values more than to 2 in input for Height and Width. Unfortunately we use fmt.ScanLn here, so it is not exactly in line with conditions of the task
//
// 	for {
// 		utils.Print(enum.EnterHnW)
// 		_, err := fmt.Scanf("%d %d\n", &h, &w)
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 			// Consume the invalid input
// 			var garbage string
// 			fmt.Scanln(&garbage)
// 			continue
// 		}

// 		if h <= 0 || w <= 0 {
// 			utils.Print(enum.ErrorHnW)
// 			continue
// 		}

// 		break
// 	}

// 	for i := 0; i < h; i++ {
// 		for j := 0; j < w; j++ {
// 			ap.PutRune('*')
// 		}
// 		ap.PutRune('\n')
// 	}
//
