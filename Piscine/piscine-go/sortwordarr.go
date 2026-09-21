package piscine

// SortWordArr sorts a string slice based on ASCII values in ascending order
func SortWordArr(a []string) {
	// Bubble Sort: Repeat the process until the list is sorted
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				// Swap if the current string has a larger ASCII value than the next string
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
