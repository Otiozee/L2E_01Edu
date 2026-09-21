package piscine

/* ListAt returns the node at the specified position in the linked list
   Parameters:
   l   - pointer to the head node of the linked list
   pos - zero-based index of the node to retrieve
   Returns:
   pointer to the node at position pos, or nil if position is invalid*/

func ListAt(l *NodeL, pos int) *NodeL {
	current := l // Start at the head of the list

	count := 0 // Initialize position counter

	for current != nil { // Traverse the list until reaching the end

		if count == pos { // Check if current position matches the requested position
			return current // Return the found node
		}

		// Move to next node and increment position counter
		count++
		current = current.Next
	}
	// Return nil if position is out of bounds
	return nil
}
