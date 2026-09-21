package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}
	if len(s) < len(toFind) {
		return -1
	}
	if s[:len(toFind)] == toFind {
		return 0
	}
	result := Index(s[1:], toFind)
	if result == -1 {
		return -1
	}
	return result + 1
}
