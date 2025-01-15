package stack

func (s *Stack) Len() int {
	count := 0
	current := s.top

	for current != nil {
		count++
		current = current.next
	}
	return count
}
