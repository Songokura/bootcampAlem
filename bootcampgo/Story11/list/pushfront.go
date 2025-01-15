package list

func (l *List) PushFront(v interface{}) {
	newNode := &ListNode{
		Value: v,
	}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
}
