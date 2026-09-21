package piscine

// Helper function to check if the base is valid
func isValidBase(base string) bool {
	// Base must have at least 2 characters
	if len(base) < 2 {
		return false
	}

	// Base should not contain '+' or '-' characters
	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			return false
		}
	}

	// Each character in the base must be unique
	for i := 0; i < len(base); i++ {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}

	return true
}

// Helper function to find the index of a rune in the base string
func indexInbase(r rune, base string) int {
	for i, br := range base {
		if br == r {
			return i
		}
	}
	return -1
}

// Function to convert a numeric string `s` in a given base `base` to an integer
func AtoiBase(s string, base string) int {
	if !isValidBase(base) {
		return 0
	}

	baseLen := len(base)
	result := 0
	for _, r := range s {
		digit := indexInbase(r, base)
		if digit == -1 {
			return 0
		}
		result = result*baseLen + digit
	}
	return result
}
