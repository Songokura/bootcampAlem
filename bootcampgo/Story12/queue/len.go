package queue

func (q *Queue) Len() int {
	count := 0
	current := q.head

	for current != nil {
		count++
		current = current.next
	}
	return count
}
