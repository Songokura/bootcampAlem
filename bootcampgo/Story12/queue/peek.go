package queue

func (q *Queue) Peek() interface{} {
	if q.head == nil {
		return nil
	}
	return q.head.value
}
