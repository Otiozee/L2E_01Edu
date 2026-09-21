package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	var word []rune
	for _, char := range s {
		if char == ' ' || char == '\t' || char == '\n' {
			if len(word) > 0 {
				result = append(result, string(word))
				word = []rune{}
			}
		} else {
			word = append(word, char)
		}
	}
	if len(word) > 0 {
		result = append(result, string(word))
	}
	return result
}
