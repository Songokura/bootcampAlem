package list

func (l *List) Remove(n *ListNode) {
	if n == nil {
		return
	}

	if n == l.Head {
		l.Head = n.Next
		if l.Head == nil {
			l.Tail = nil
		}
		return
	}

	current := l.Head
	for current != nil && current.Next != n {
		current = current.Next
	}

	if current != nil {
		current.Next = n.Next
		if n == l.Tail {
			l.Tail = current
		}
	}
}
