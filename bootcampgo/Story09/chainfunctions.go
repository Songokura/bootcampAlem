package bootcamp

func ChainFunctions(funcs []func(*string)) func(*string) {
	return func(str *string) {
		for _, f := range funcs {
			f(str)
		}
	}
}

// func main() {
// 	toUpper := func(str *string) {
// 		*str = strings.ToUpper(*str)
// 	}

// 	addExclamation := func(str *string) {
// 		*str = *str + "!"
// 	}

// 	reverseStr := func(str *string) {
// 		runes := []rune(*str)
// 		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 			runes[i], runes[j] = runes[j], runes[i]
// 		}
// 		*str = string(runes)
// 	}

// 	chainedFunc := ChainFunctions([]func(*string){
// 		toUpper,
// 		addExclamation,
// 		reverseStr,
// 	})

// 	str := "salem"
// 	chainedFunc(&str)
// 	fmt.Println(str) // "!MELAS"
// }
