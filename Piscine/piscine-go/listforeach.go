package piscine

// ListForEach applies a function to each node in the list
func ListForEach(l *List, f func(*NodeL)) {
	if l == nil || l.Head == nil {
		return
	}

	current := l.Head
	for current != nil {
		f(current)             // Apply the function to the current node
		current = current.Next // Move to next node
	}
}

// Add2_node adds 2 to int nodes or appends "2" to string nodes
func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

// Subtract3_node subtracts 3 from int nodes or appends "-3" to string nodes
func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}
