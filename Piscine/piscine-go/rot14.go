package piscine

func Rot14(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		if r >= 'a' && r <= 'z' {
			// Shift within lowercase letters
			runes[i] = 'a' + ((r - 'a' + 14) % 26)
		} else if r >= 'A' && r <= 'Z' {
			// Shift within uppercase letters
			runes[i] = 'A' + ((r - 'A' + 14) % 26)
		}
		// Other characters stay the same
	}
	return string(runes)
}
