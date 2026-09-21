package piscine

// IsPositiveNode checks if node data is a positive number
func IsPositiveNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int:
		return node.Data.(int) > 0
	case float32:
		return node.Data.(float32) > 0
	case float64:
		return node.Data.(float64) > 0
	case byte:
		return node.Data.(byte) > 0
	default:
		return false
	}
}

// IsAlNode checks if node data is not a number
func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}

// ListForEachIf applies function f to nodes that satisfy condition cond
func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	if l == nil || l.Head == nil {
		return
	}

	current := l.Head
	for current != nil {
		if cond(current) {
			f(current)
		}
		current = current.Next
	}
}

// StringToInt converts string nodes to integer value 2
func StringToInt(node *NodeL) {
	node.Data = 2
}
