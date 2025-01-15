package bootcamp

func IsNil(a *int) bool {
	return a == nil
}

// func main() {
//	var a int = 10
//	var b *int
//	fmt.Println(a, IsNil(&a))    // 10 false
//	fmt.Println(b, IsNil(b))     // <nil> true
//	fmt.Println(nil, IsNil(nil)) // <nil> true
//}
