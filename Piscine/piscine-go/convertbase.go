package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	if len(baseFrom) < 2 || len(baseTo) < 2 {
		return ""
	}

	// Convert nbr from baseFrom to decimal (base 10)
	decimalValue := 0
	baseFromLen := len(baseFrom)
	for _, char := range nbr {
		digitValue := 0
		for j, b := range baseFrom {
			if b == char {
				digitValue = j
				break
			}
		}
		decimalValue = decimalValue*baseFromLen + digitValue
	}

	// Handle the special case where decimalValue is 0
	if decimalValue == 0 {
		return string(baseTo[0])
	}

	// Convert decimalValue to baseTo
	var result []rune
	baseToLen := len(baseTo)
	for decimalValue > 0 {
		remainder := decimalValue % baseToLen
		result = append(result, rune(baseTo[remainder]))
		decimalValue = decimalValue / baseToLen
	}

	// Reverse the result slice
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}
