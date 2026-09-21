package piscine

func ListClear(l *List) {
	// If the list pointer is nil, there's nothing to clear
	if l == nil {
		return
	}
	// Set Head pointer to nil, breaking the link to the first node
	l.Head = nil
	// Set Tail pointer to nil, breaking the link to the last node
	l.Tail = nil
}
