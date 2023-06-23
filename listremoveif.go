package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	if l.Head == nil {
		l.Tail = nil
		return
	}

	prev := l.Head
	curr := l.Head.Next

	for curr != nil {
		if curr.Data == data_ref {
			prev.Next = curr.Next
			curr = curr.Next
		} else {
			prev = curr
			curr = curr.Next
		}
	}
}
