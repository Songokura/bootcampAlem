package bootcamp

// import "fmt"

func Concat(s ...string) string {
	var concated string
	for _, v := range s {
		concated += v
	}
	return concated
}

// func main() {
// 	fmt.Println(Concat("Salem", " ", "Student!"))
// 	fmt.Println(Concat("1", "2", "3", "4", "5"))
// }
