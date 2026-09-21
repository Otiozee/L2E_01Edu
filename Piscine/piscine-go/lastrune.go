package piscine

func LastRune(s string) rune {
	if len(s) == 0 {
		return 0 // Returns zero if the string is empty
	}
	return []rune(s)[len(s)-1] // Convert string to rune slice and return the last rune
}
