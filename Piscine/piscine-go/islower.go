package piscine

// Omo this islower func na to checks if the string contains only lowercase characters by comparing lengths.
func IsLower(s string) bool {
	upperCount := 0
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			upperCount++
		}
	}
	return upperCount == len(s)
}
