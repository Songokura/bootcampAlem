package stack

type Node struct {
	value interface{}
	next  *Node
}

type Stack struct {
	top *Node
}

func NewStack() *Stack {
	return &Stack{
		top: nil,
	}
}

func (s *Stack) Push(v interface{}) {
	newNode := &Node{
		value: v,
		next:  s.top,
	}
	s.top = newNode
}

func (s *Stack) Pop() interface{} {
	if s.top == nil {
		return nil
	}
	poppedValue := s.top.value
	s.top = s.top.next
	return poppedValue
}
