package piscine

// AdvancedSortWordArr sorts the slice of strings based on the comparison function f
func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// Use the comparison function to decide whether to swap
			if f(a[j], a[j+1]) > 0 {
				// Swap a[j] and a[j+1] if they are out of order
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
