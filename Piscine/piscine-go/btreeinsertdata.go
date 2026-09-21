package piscine

// TreeNode represents a node in a binary search tree
// with left and right child pointers and a parent pointer
type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// BTreeInsertData inserts a new node with the given data into the binary search tree
// while maintaining the BST properties (left < parent < right)
//
// Parameters:
//
//	root - the root node of the tree (or subtree) where insertion should occur
//	data - the string value to be inserted into the tree
//
// Returns:
//
//	*TreeNode - the root node of the modified tree
//
// Notes:
// - Duplicate values are ignored (not inserted)
// - Maintains parent pointers for all nodes
// - Uses recursion to find the correct insertion point
func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	// Base case: if current node is nil, create a new node
	if root == nil {
		return &TreeNode{Data: data}
	}

	// Insert into left subtree if data is smaller than current node's data
	if data < root.Data {
		if root.Left == nil {
			// Create new left child with current node as parent
			root.Left = &TreeNode{Data: data, Parent: root}
		} else {
			// Recursively insert into left subtree
			BTreeInsertData(root.Left, data)
		}
	} else if data > root.Data {
		// Insert into right subtree if data is larger than current node's data
		if root.Right == nil {
			// Create new right child with current node as parent
			root.Right = &TreeNode{Data: data, Parent: root}
		} else {
			// Recursively insert into right subtree
			BTreeInsertData(root.Right, data)
		}
	}
	// Note: if data == root.Data, we do nothing (no duplicates)

	// Return the root of the (sub)tree
	return root
}
