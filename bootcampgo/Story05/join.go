package bootcamp

// import "fmt"

func Join(elements []string, sep string) string {
	if len(elements) == 0 {
		return ""
	}

	result := elements[0]
	for _, v := range elements[1:] {
		result += sep + v
	}
	return result
}

// func main() {
// 	fmt.Println(Join([]string{"salem", "student"}, " "))
// 	fmt.Println(Join([]string{"1", "2", "3"}, ", "))
// }
