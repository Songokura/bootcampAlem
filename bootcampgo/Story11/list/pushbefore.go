package list

func (l *List) PushBefore(n *ListNode, v interface{}) {
	if n == nil {
		return
	}
	newNode := &ListNode{
		Value: v,
	}

	if n == l.Head {
		newNode.Next = l.Head
		l.Head = newNode
		return
	}

	current := l.Head
	for current != nil && current.Next != n {
		current = current.Next
	}

	if current == nil {
		return
	}
	current.Next = newNode
	newNode.Next = n
}
