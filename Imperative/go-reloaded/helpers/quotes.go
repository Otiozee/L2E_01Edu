package helpers

import "regexp"

// FixQuotes removes extra spaces inside single quotes.
// Example:
// ' hello ' to 'hello'
func FixQuotes(s string) string {

	// This regex finds text inside single quotes with possible spaces
	// and keeps only the inner content without extra spaces
	re := regexp.MustCompile(`'\s*(.*?)\s*'`)

	// Replace with cleaned version: remove spaces inside quotes
	return re.ReplaceAllString(s, "'$1'")
}
