package piscine

func Compare(a, b string) int {
	minLength := len(a)
	if len(b) < minLength {
		minLength = len(b)
	}

	for i := 0; i < minLength; i++ {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}
	if len(a) < len(b) {
		return -1
	} else if len(a) > len(b) {
		return 1
	}
	return 0
}
