package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	// Create a dummy node to build the merged list
	dummy := &NodeI{}
	current := dummy

	// Traverse both lists
	for n1 != nil && n2 != nil {
		if n1.Data <= n2.Data {
			current.Next = n1
			n1 = n1.Next
		} else {
			current.Next = n2
			n2 = n2.Next
		}
		current = current.Next
	}

	// Attach the remaining elements from either list
	if n1 != nil {
		current.Next = n1
	} else {
		current.Next = n2
	}

	// Return the merged list (skip the dummy node)
	return dummy.Next
}
