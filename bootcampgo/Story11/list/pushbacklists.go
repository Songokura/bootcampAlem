package list

func (l *List) PushBackLists(lists ...*List) {
	for _, list := range lists {
		if list == nil || list.Head == nil {
			continue
		}
		if l.Head == nil {
			l.Head = list.Head
			l.Tail = list.Tail
		} else {
			l.Tail.Next = list.Head
			l.Tail = list.Tail
		}
	}
}
