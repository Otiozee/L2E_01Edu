package piscine

func Split(s, sep string) []string {
	var result []string
	start := 0
	sepLen := len(sep)
	if sepLen == 0 {
		for _, char := range s {
			result = append(result, string(char))
		}
		return result
	}
	for i := 0; i <= len(s)-sepLen; i++ {
		if s[i:i+sepLen] == sep {
			result = append(result, s[start:i])
			start = i + sepLen
			i += sepLen - 1
		}
	}
	result = append(result, s[start:])
	return result
}
