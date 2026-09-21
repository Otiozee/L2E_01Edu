package piscine

func IsNumeric(s string) bool {
	if s == "" {
		return true
	}
	for _, char := range s {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
