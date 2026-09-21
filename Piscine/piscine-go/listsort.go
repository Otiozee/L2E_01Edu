package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

// ListSort sorts a linked list in ascending order using merge sort
func ListSort(l *NodeI) *NodeI {
	// Base case: empty list or single element
	if l == nil || l.Next == nil {
		return l
	}

	// Split the list into two halves
	mid := getMiddle(l)
	left := l
	right := mid.Next
	mid.Next = nil

	// Recursively sort both halves
	left = ListSort(left)
	right = ListSort(right)

	// Merge the sorted halves
	return merge(left, right)
}

// getMiddle finds the middle node of a linked list
func getMiddle(head *NodeI) *NodeI {
	if head == nil {
		return head
	}

	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// merge merges two sorted linked lists
func merge(left, right *NodeI) *NodeI {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	var result *NodeI

	if left.Data <= right.Data {
		result = left
		result.Next = merge(left.Next, right)
	} else {
		result = right
		result.Next = merge(left, right.Next)
	}

	return result
}
