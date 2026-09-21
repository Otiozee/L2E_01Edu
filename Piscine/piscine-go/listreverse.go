package piscine

func ListReverse(l *List) {
	if l == nil || l.Head == nil {
		return
	}

	var previous *NodeL
	current := l.Head
	l.Tail = l.Head // The current head will become the tail after reversal

	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}

	l.Head = previous // The last node becomes the new head
}
