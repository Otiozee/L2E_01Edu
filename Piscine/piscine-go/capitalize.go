package piscine

func Capitalize(s string) string {
	var result string
	inWord := false

	for i := 0; i < len(s); i++ {
		char := s[i]

		// Check if the character is alphanumeric (A-Z, a-z, 0-9)
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
			// If it's the start of a new word, capitalize it
			if !inWord {
				// Capitalize the first letter
				if char >= 'a' && char <= 'z' {
					char -= 'a' - 'A' // Convert to uppercase
				}
				inWord = true
			} else {
				// Lowercase the rest of the letters in the word
				if char >= 'A' && char <= 'Z' {
					char += 'a' - 'A' // Convert to lowercase
				}
			}
		} else {
			// Non-alphanumeric characters are added as is
			inWord = false
		}

		// Add the processed character to the result string
		result += string(char)
	}

	return result
}
