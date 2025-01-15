package list

func (l *List) PushAfter(n *ListNode, v interface{}) {
	if n == nil {
		return
	}

	newNode := &ListNode{
		Next:  n.Next,
		Value: v,
	}
	n.Next = newNode
	if l.Tail == n {
		l.Tail = newNode
	}
}
