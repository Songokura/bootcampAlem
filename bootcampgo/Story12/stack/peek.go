package stack

func (s *Stack) Peek() interface{} {
	if s.top == nil {
		return nil
	}

	return s.top.value
}
