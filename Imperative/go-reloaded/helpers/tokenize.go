package helpers

import "regexp"

// Tokenize splits a raw string into smaller pieces (tokens).
// These tokens include words, punctuation, quotes, and special patterns.
func Tokenize(s string) []string {

	// This regex breaks text into meaningful parts:
	// \n        - newline
	// ...       - ellipsis
	// [!?]+     - multiple exclamation/question marks
	// [.,:;!?'] - punctuation marks and single quote
	// \([^)]+\) - anything inside parentheses (like (hex), (up,2))
	// \S+       - normal words (non-space characters)
	re := regexp.MustCompile(`\n|\.\.\.|[!?]+|[.,:;!?']|\([^)]+\)|\S+`)

	// Find all matching tokens in the input string
	return re.FindAllString(s, -1)
}
