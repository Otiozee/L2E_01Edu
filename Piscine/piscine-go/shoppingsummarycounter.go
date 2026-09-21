package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	start := 0

	for i, char := range str {
		if char == ' ' {
			item := str[start:i]
			summary[item]++
			start = i + 1
		}
	}
	// Add the last segment
	item := str[start:]
	summary[item]++

	return summary
}
