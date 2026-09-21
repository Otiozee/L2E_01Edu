package helpers

import "strings"

// Rebuild converts a list of tokens back into a normal string sentence.
// It also handles spacing and new lines properly.
func Rebuild(tokens []string) string {

	// strings.Builder is used for efficient string building
	var b strings.Builder

	// loop through all tokens one by one
	for i, t := range tokens {

		// if token is a newline, add it directly and continue
		if t == "\n" {
			b.WriteString("\n")
			continue
		}

		// add the current token to the result
		b.WriteString(t)

		// decide whether to add a space after this token
		if i < len(tokens)-1 {

			next := tokens[i+1]

			// do not add space before a newline
			if next != "\n" {
				b.WriteString(" ")
			}
		}
	}

	// return final reconstructed string
	return b.String()
}
