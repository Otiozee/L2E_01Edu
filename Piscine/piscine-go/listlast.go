package piscine

/*
ListLast returns the data of the last element in the list
If the list is empty or nil, it returns nil
*/
func ListLast(l *List) interface{} {
	// Check if the list is nil or the Tail pointer is nil (empty list)
	if l == nil || l.Tail == nil {
		return nil
	}
	// Return the data of the Tail node (last element)
	return l.Tail.Data
}
