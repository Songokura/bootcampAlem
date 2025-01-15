package list

func (l *List) Len() int {
	count := 0
	current := l.Head

	for current != nil {
		count++
		current = current.Next
	}
	return count
}
