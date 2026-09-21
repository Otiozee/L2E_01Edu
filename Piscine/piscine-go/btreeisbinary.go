package piscine

func BTreeIsBinary(root *TreeNode) bool {
	var check func(node *TreeNode, min, max string) bool
	check = func(node *TreeNode, min, max string) bool {
		if node == nil {
			return true
		}

		if min != "" && node.Data <= min {
			return false
		}
		if max != "" && node.Data >= max {
			return false
		}

		return check(node.Left, min, node.Data) &&
			check(node.Right, node.Data, max)
	}

	return check(root, "", "")
}
