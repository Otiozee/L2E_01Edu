package piscine

// NodeL represents a node in a linked list
type NodeL struct {
	Data interface{} // Data can hold any type of value
	Next *NodeL      // Next points to the next node in the list
}

// List represents a linked list with Head and Tail pointers
type List struct {
	Head *NodeL // Head points to the first node in the list
	Tail *NodeL // Tail points to the last node in the list
}

// ListPushBack inserts a new node with given data at the end of the list
func ListPushBack(l *List, data interface{}) {
	// Create a new node containing the data
	newNode := &NodeL{Data: data}

	// If list is empty, set both Head and Tail to the new node
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		// If list is not empty:
		// 1. Link current Tail node to new node
		l.Tail.Next = newNode
		// 2. Update Tail to point to the new last node
		l.Tail = newNode
	}
}
