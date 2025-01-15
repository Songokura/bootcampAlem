package list

type ListNode struct {
	Next  *ListNode
	Prev  *ListNode
	Value interface{}
}

type List struct {
	Head *ListNode
	Tail *ListNode
}

func NewList() *List {
	return &List{
		Head: nil,
		Tail: nil,
	}
}

// func main() {
// 	l := NewList()
// 	fmt.Println(l)
// 	// Output: &{<nil> <nil>}
// }
