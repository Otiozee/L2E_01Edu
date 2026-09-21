package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) <= 1 {
		return true
	}

	// Determine direction from first non-equal comparison
	direction := 0
	for i := 0; i < len(a)-1 && direction == 0; i++ {
		direction = f(a[i], a[i+1])
	}

	// Check all elements follow the determined direction
	for i := 0; i < len(a)-1; i++ {
		comp := f(a[i], a[i+1])
		if direction > 0 { // Expect descending
			if comp < 0 {
				return false
			}
		} else if direction < 0 { // Expect ascending
			if comp > 0 {
				return false
			}
		}
		// If direction == 0 (all equal), any comp is fine
	}

	return true
}
