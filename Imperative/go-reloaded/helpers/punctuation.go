package helpers

import (
	"regexp"
	"strings"
)

// FixPunctuation cleans and formats punctuation in a string.
// It removes extra spaces around punctuation and fixes quotes spacing.
func FixPunctuation(s string) string {

	// This regex finds punctuation like ., : ; ! ? with optional spaces around them
	// Example: "hello ?" to "hello?"
	re := regexp.MustCompile(`[ ]*([.,:;!?]+)[ ]*`)
	s = re.ReplaceAllString(s, "$1 ")

	// This fixes spacing inside double quotes
	// Example: " Pot ? " to "Pot?"
	reQuotes := regexp.MustCompile(`"\s*([^"]*?)\s*"`)
	s = reQuotes.ReplaceAllString(s, `"$1"`)

	// Remove extra space at the end of the string
	return strings.TrimRight(s, " ")
}
