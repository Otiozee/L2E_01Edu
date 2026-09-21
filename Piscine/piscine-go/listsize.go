package piscine

func ListSize(l *List) int {
	// Initialize a counter to zero
	count := 0
	// Start at the head of the list
	current := l.Head
	// Traverse the list until reaching the end (nil)
	for current != nil {
		// Increment count for each node encountered
		count++
		// Move to the next node
		current = current.Next
	}
	// Return total number of nodes counted
	return count
}
