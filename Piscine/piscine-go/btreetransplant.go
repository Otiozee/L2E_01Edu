package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node == nil {
		return root
	}

	// If node is the root of the entire tree
	if node.Parent == nil {
		// Set rplc as new root
		if rplc != nil {
			rplc.Parent = nil
		}
		return rplc
	}

	// Update parent's reference to the node
	if node == node.Parent.Left {
		node.Parent.Left = rplc
	} else {
		node.Parent.Right = rplc
	}

	// Update rplc's parent pointer
	if rplc != nil {
		rplc.Parent = node.Parent
	}

	return root
}
