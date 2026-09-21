package piscine

// Omo this isUpper func na to checks if the string contains only uppercase characters by comparing lengths.
func IsUpper(s string) bool {
	upperCount := 0
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			upperCount++
		}
	}
	return upperCount == len(s)
}
