package piscine

// The following command concatenates the elements of the slice strs using the separator sep.
func Join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	return strs[0] + sep + Join(strs[1:], sep)
}
