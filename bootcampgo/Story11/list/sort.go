package list

func (l *List) Sort(fn func(a *ListNode, b *ListNode) int) {
	if l.Head == nil || l.Head == l.Tail {
		return
	}

	swapped := true

	for swapped {
		swapped = false
		current := l.Head
		for current != nil && current.Next != nil {
			next := current.Next
			if fn(current, next) > 0 {
				current.Value, next.Value = next.Value, current.Value
				swapped = true
			}
			current = current.Next
		}
	}
}
