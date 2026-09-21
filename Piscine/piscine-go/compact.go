package piscine

func Compact(ptr *[]string) int {
	slice := *ptr
	count := 0
	for _, v := range slice {
		if v != "" {
			count++
		}
	}
	compacted := make([]string, 0, count)
	for _, v := range slice {
		if v != "" {
			compacted = append(compacted, v)
		}
	}
	*ptr = compacted
	return count
}
