package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	// Handle empty list case
	if l == nil || l.Head == nil {
		return
	}

	// Remove matching nodes from the head
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
		if l.Head == nil {
			l.Tail = nil
		}
	}

	// If all nodes were removed
	if l.Head == nil {
		return
	}

	// Remove matching nodes after the head
	current := l.Head
	for current.Next != nil {
		if current.Next.Data == data_ref {
			current.Next = current.Next.Next
			// Update tail if we removed the last node
			if current.Next == nil {
				l.Tail = current
			}
		} else {
			current = current.Next
		}
	}
}
