package queue

type Node struct {
	value interface{}
	next  *Node
}

type Queue struct {
	head *Node
	tail *Node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(v interface{}) {
	newNode := &Node{value: v}

	if q.tail == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
}

func (q *Queue) Pop() interface{} {
	if q.head == nil {
		return nil
	}
	value := q.head.value
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}

	return value
}
