package piscine

// CompStr compares two interface values for equality using the == operator.
// This simple comparison works for basic types like strings, numbers, and other comparable types.
//
// Parameters:
//
//	a - The first value to compare (interface{})
//	b - The second value to compare (interface{})
//
// Returns:
//
//	bool - true if the values are equal, false otherwise
func CompStr(a, b interface{}) bool {
	return a == b
}

// ListFind searches through a linked list to find the first node containing data
// that matches the reference value according to the provided comparison function.
//
// Parameters:
//
//	l    - Pointer to the List structure to search through
//	ref  - The reference value to compare against node data (interface{})
//	comp - The comparison function that determines equality between two values
//
// Returns:
//
//	*interface{} - Pointer to the matching data if found, nil otherwise
//
// Notes:
// - The function performs a linear search through the list (O(n) time complexity)
// - Returns the first match found
// - Safe to call with nil list (returns nil)
func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	// Start traversal at the head of the list
	current := l.Head

	// Iterate through each node until the end of the list
	for current != nil {
		// Use the provided comparison function to check for equality
		if comp(current.Data, ref) {
			// Return the address of the matching data
			return &current.Data
		}
		// Move to the next node
		current = current.Next
	}

	// Return nil if no matching node was found
	return nil
}
