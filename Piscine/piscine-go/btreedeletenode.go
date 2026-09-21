package piscine

// BTreeDeleteNode removes a node from the BST while maintaining BST properties
func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}

	// Case 1: Node has no children
	if node.Left == nil && node.Right == nil {
		if node.Parent == nil {
			return nil // Deleting the root node
		}
		if node == node.Parent.Left {
			node.Parent.Left = nil
		} else {
			node.Parent.Right = nil
		}
		return root
	}

	// Case 2: Node has one child
	if node.Left == nil {
		return BTreeTransplant(root, node, node.Right)
	}
	if node.Right == nil {
		return BTreeTransplant(root, node, node.Left)
	}

	// Case 3: Node has two children
	successor := BTreeMin(node.Right)
	if successor.Parent != node {
		root = BTreeTransplant(root, successor, successor.Right)
		successor.Right = node.Right
		successor.Right.Parent = successor
	}
	root = BTreeTransplant(root, node, successor)
	successor.Left = node.Left
	successor.Left.Parent = successor

	return root
}
