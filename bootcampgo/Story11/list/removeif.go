package list

func (l *List) RemoveIf(fn func(n *ListNode) bool) {
	var prev *ListNode
	current := l.Head

	for current != nil {
		if fn(current) {
			if prev != nil {
				prev.Next = current.Next
				if current == l.Tail {
					l.Tail = prev
				}
			} else {
				l.Head = current.Next
				if current == l.Tail {
					l.Tail = nil
				}
			}
		} else {
			prev = current
		}
		current = current.Next
	}
}
