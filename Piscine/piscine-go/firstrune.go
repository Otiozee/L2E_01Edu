package piscine

func FirstRune(s string) rune {
	if len(s) == 0 {
		return 0 // This will return zero value for empty string
	}
	runes := []rune(s) // This will convert string to rune slice
	return runes[0]    // This will return the first rune
}
