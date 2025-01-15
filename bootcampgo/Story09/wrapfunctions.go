package bootcamp

import (
	"github.com/alem-platform/ap"
)

func WrapperPrintStr(fn func(str *string)) func(str *string) {
	return func(str *string) {
		for _, v := range *str {
			ap.PutRune(v)
		}
		ap.PutRune('\n')
		fn(str)
	}
}

func WrapperRot1(fn func(str *string)) func(str *string) {
	return func(str *string) {
		runes := []rune(*str)
		for i := 0; i < len(runes); i++ {
			switch {
			case 'a' <= runes[i] && runes[i] <= 'z':
				runes[i] = 'a' + ((runes[i] - 'a' + 1) % 26)
			case 'A' <= runes[i] && runes[i] <= 'Z':
				runes[i] = 'A' + ((runes[i] - 'A' + 1) % 26)
			}
		}
		*str = string(runes)
		fn(str)
	}
}

func WrapperRot13(fn func(str *string)) func(str *string) {
	return func(str *string) {
		runes := []rune(*str)
		for i := 0; i < len(runes); i++ {
			switch {
			case 'a' <= runes[i] && runes[i] <= 'z':
				runes[i] = 'a' + ((runes[i] - 'a' + 13) % 26)
			case 'A' <= runes[i] && runes[i] <= 'Z':
				runes[i] = 'A' + ((runes[i] - 'A' + 13) % 26)
			}
		}
		*str = string(runes)
		fn(str)
	}
}

func WrapperReverseStr(fn func(str *string)) func(str *string) {
	return func(str *string) {
		runes := []rune(*str)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		*str = string(runes)
		fn(str)
	}
}

func WrapFunctions(decs []func(fn func(str *string)) func(str *string)) func(str *string) {
	return func(str *string) {
		for _, v := range decs {
			fn := v(func(s *string) {})
			fn(str)
		}
	}
}

// func main() {
// 	mockFn := func(str *string) {
// 		return
// 	}

// 	fnPrint := WrapperPrintStr(mockFn)

// 	str := "salem?"
// 	fnPrint(&str) // salem?

// 	fnRot1 := WrapperRot1(mockFn)
// 	fnRot1(&str)
// 	fnPrint(&str) // tbmfn?

// 	fnRot13 := WrapperRot13(mockFn)
// 	fnRot13(&str)
// 	fnPrint(&str) // gozsa?

// 	fnReverse := WrapperReverseStr(mockFn)
// 	fnReverse(&str)
// 	fnPrint(&str) // ?aszog

// 	fmt.Println("United Func Results")
// 	wrappedFns := WrapFunctions([]func(fn func(str *string)) func(str *string){
// 		WrapperPrintStr,
// 		WrapperRot1,
// 		WrapperRot1,
// 		WrapperRot1,
// 		WrapperRot1,
// 		WrapperRot1,
// 		WrapperRot1,
// 		WrapperRot1,
// 		WrapperRot1,
// 		WrapperRot1,
// 		WrapperRot1,
// 		WrapperRot1,
// 		WrapperRot1,
// 		WrapperPrintStr,
// 		WrapperReverseStr,
// 		WrapperPrintStr,
// 	})
// 	wrappedFns(&str)
// 	// ?aszog
// 	// gozsa?
// 	// salem?
// }
