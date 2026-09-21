package piscine

// ListPushFront inserts a new node with given data at the beginning of the list
func ListPushFront(l *List, data interface{}) {
	// Create a new node containing the data
	newNode := &NodeL{Data: data}

	// If list is empty, set both Head and Tail to the new node
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		/*If list is not empty:
		1. Link current head node to new node*/
		newNode.Next = l.Head

		// 2. Update head to point to the new 1st node
		l.Head = newNode
	}
}
