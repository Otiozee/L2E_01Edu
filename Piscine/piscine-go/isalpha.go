package piscine

func IsAlpha(s string) bool {
	if s == "" {
		return true
	}
	for _, char := range s {
		if !((char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')) {
			return false
		}
	}
	return true
}
