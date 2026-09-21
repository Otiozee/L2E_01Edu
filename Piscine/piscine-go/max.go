package piscine

func Max(a []int) int {
	// Check if the slice is empty
	if len(a) == 0 {
		return 0
	}

	// Initialize the max variable to the first element of the slice
	max := a[0]

	// Loop through the slice to find the maximum value
	for _, value := range a {
		if value > max {
			max = value
		}
	}

	return max
}
