package piscine

func BasicAtoi(s string) int {
	result := 0
	for _, char := range s {
		result = result*10 + int(char-'0')
	}
	return result
}
