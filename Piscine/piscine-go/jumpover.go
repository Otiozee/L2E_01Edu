package piscine

func JumpOver(str string) string {
	result := ""
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}
	if len(result) == 0 {
		return "\n"
	}
	return result + "\n"
}
