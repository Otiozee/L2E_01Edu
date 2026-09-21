package helpers

import (
	"strconv"
	"strings"
)

// ApplyModifiers processes special commands inside tokens like:
// (hex), (bin), (up), (low), (cap)
//
// It modifies previous words based on these commands.
func ApplyModifiers(tokens []string) []string {

	// loop through all tokens
	for i := 0; i < len(tokens); i++ {

		t := tokens[i]

		// check if token is a modifier like "(hex)" or "(up,2)"
		if strings.HasPrefix(t, "(") && strings.HasSuffix(t, ")") {

			// remove parentheses: "(hex,2)" to "hex,2"
			content := strings.Trim(t, "()")

			// split command and optional number: ["hex", "2"]
			parts := strings.Split(content, ",")

			// first part is the command (hex, bin, up, low, cap)
			cmd := strings.TrimSpace(parts[0])

			// default: apply to 1 previous word
			count := 1

			// if user gives number (like 2), parse it
			if len(parts) == 2 {
				n, err := strconv.Atoi(strings.TrimSpace(parts[1]))
				if err == nil && n > 0 {
					count = n
				}
			}

			// find index of previous word(s) to modify
			start := FindPreviousWords(tokens, i, count)

			// apply the correct modifier
			switch cmd {

			// convert hexadecimal to decimal
			case "hex":
				if start >= 0 {
					if val, err := strconv.ParseInt(tokens[start], 16, 64); err == nil {
						tokens[start] = strconv.FormatInt(val, 10)
					}
				}

			// convert binary to decimal
			case "bin":
				if start >= 0 {
					if val, err := strconv.ParseInt(tokens[start], 2, 64); err == nil {
						tokens[start] = strconv.FormatInt(val, 10)
					}
				}

			// convert previous words to uppercase
			case "up":
				ApplyCase(tokens, start, i, strings.ToUpper)

			// convert previous words to lowercase
			case "low":
				ApplyCase(tokens, start, i, strings.ToLower)

			// capitalize previous words
			case "cap":
				ApplyCase(tokens, start, i, Capitalize)
			}

			// remove the modifier token after applying it
			tokens = append(tokens[:i], tokens[i+1:]...)

			// move index back because we removed an element
			i--
		}
	}

	// return modified token list
	return tokens
}
