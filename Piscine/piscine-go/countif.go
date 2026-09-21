package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0 // Initialize a counter to keep track of how many elements satisfy the condition

	// Loop through each element of the slice 'tab'
	for _, v := range tab {
		if f(v) { // Apply function 'f' on the string 'v'
			count++ // If 'f(v)' is true, increment the count
		}
	}

	return count // Return the final count
}
