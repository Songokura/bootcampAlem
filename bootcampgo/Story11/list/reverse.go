package list

func (l *List) Reverse() {
	var Prev *ListNode

	current := l.Head
	l.Tail = l.Head

	for current != nil {
		next := current.Next
		current.Next = Prev
		current.Prev = next
		Prev = current
		current = next
	}
	l.Head = Prev
}
