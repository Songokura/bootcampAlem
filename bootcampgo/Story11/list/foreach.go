package list

func (l *List) ForEach(fn func(n *ListNode)) {
	current := l.Head

	for current != nil {
		fn(current)
		current = current.Next
	}
}
