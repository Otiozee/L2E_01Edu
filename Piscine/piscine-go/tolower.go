package piscine

func ToLower(s string) string {
	runes := []rune(s)
	for i, char := range runes {
		if char >= 'A' && char <= 'Z' {
			runes[i] = char + 32 // na dis wan go convert uppercase to lowercase
		}
	}
	return string(runes)
}
