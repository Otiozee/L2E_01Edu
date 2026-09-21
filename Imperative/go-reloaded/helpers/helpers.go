package helpers

import (
	"regexp"
	"strings"
	"unicode"
)

// FindPreviousWords searches backwards in the token list
// and returns the index of the Nth previous word.
// Example:
// tokens = ["I", "love", "go"]
// if count = 1 from "go", it returns index of "love"
func FindPreviousWords(tokens []string, pos, count int) int {

	found := 0

	// loop backwards from current position
	for i := pos - 1; i >= 0; i-- {

		// check if current token is a valid word
		if IsWord(tokens[i]) {
			found++

			// when we reach required count, return index
			if found == count {
				return i
			}
		}
	}

	// return -1 if not found
	return -1
}

// ApplyCase applies a transformation function (fn)
// to words between start and end positions.
func ApplyCase(tokens []string, start, end int, fn func(string) string) {

	// safety check: if start is invalid, do nothing
	if start < 0 {
		return
	}

	// loop through selected range
	for i := start; i < end; i++ {

		// only modify valid words (ignore punctuation)
		if IsWord(tokens[i]) {
			tokens[i] = fn(tokens[i])
		}
	}
}

// Capitalize makes the first letter uppercase
// and the rest lowercase.
//
// Example: "hELLO" → "Hello"
func Capitalize(s string) string {

	if len(s) == 0 {
		return s
	}

	// convert string to lowercase first
	runes := []rune(strings.ToLower(s))

	// make first letter uppercase
	runes[0] = unicode.ToUpper(runes[0])

	return string(runes)
}

// IsWord checks if a token is a valid word
// (only letters and numbers, no punctuation)
func IsWord(s string) bool {

	// regex: only A-Z, a-z, 0-9 allowed
	return regexp.MustCompile(`^[A-Za-z0-9]+$`).MatchString(s)
}
