package helpers

import (
	"regexp"
	"strings"
)

// aSpecial contains words that should ALWAYS use "a"
// even if they start with vowels (because pronunciation starts with a consonant sound)
var aSpecial = map[string]bool{
	"university": true, "european": true, "one": true, "user": true,
	"useful": true, "unicorn": true, "unit": true, "eulogy": true,
	"euphemism": true, "ukulele": true, "unique": true, "once": true,
	"uniform": true, "union": true, "utensil": true, "utility": true,
	"universe": true, "utopia": true, "usb": true,
	"u-turn": true, "usa": true,
}

// anSpecial contains words that should ALWAYS use "an"
// because they start with a silent "h" or vowel sound
var anSpecial = map[string]bool{
	"hour": true, "honest": true, "honor": true, "heir": true,
	"honour": true, "honorable": true, "honourable": true,
	"hourglass": true, "heiress": true,
	"heirloom": true, "hourly": true, "fbi": true,
}

// This regex finds patterns like:
// "an apple", "a house", "An idea", etc.
// It captures the article and the word after it
var re = regexp.MustCompile(`\b([Aa][Nn]?)\s+('?[\w-]+'?)`)

// cleanText fixes spacing issues before article processing
func cleanText(s string) string {

	// Fix words inside single quotes:
	// example: ' house ' to 'house'
	reQuotes := regexp.MustCompile(`'\s*([^']*?)\s*'`)
	s = reQuotes.ReplaceAllString(s, "'$1'")

	// Remove space before punctuation:
	// example: "man ?" to "man?"
	rePunct := regexp.MustCompile(`\s+([.,!?])`)
	s = rePunct.ReplaceAllString(s, "$1")

	return s
}

// FixArticles corrects wrong usage of "a" and "an"
func FixArticles(s string) string {

	// first clean the text (fix quotes + punctuation spacing)
	s = cleanText(s)

	return re.ReplaceAllStringFunc(s, func(match string) string {

		// split match into article + word
		parts := re.FindStringSubmatch(match)
		if len(parts) != 3 {
			return match
		}

		article := parts[1]
		word := parts[2]

		// clean the word by removing quotes and punctuation
		cleanWord := word

		// remove starting and ending single quotes (clean way)
		cleanWord = strings.TrimPrefix(cleanWord, "'")
		cleanWord = strings.TrimSuffix(cleanWord, "'")

		// remove punctuation like . , ! ?
		cleanWord = strings.Trim(cleanWord, ".,!?;:")

		// get the correct article ("a" or "an")
		correctArticle := getCorrectArticle(cleanWord)

		// if already correct, keep original
		if strings.ToLower(article) == correctArticle {
			return article + " " + word
		}

		// preserve capitalization (A vs a)
		if article[0] >= 'A' && article[0] <= 'Z' {
			return strings.ToUpper(correctArticle[:1]) + correctArticle[1:] + " " + word
		}

		return correctArticle + " " + word
	})
}

// getCorrectArticle decides whether a word uses "a" or "an"
func getCorrectArticle(word string) string {

	wordLower := strings.ToLower(word)
	wordUpper := strings.ToUpper(word)

	// check special "an" words first
	if anSpecial[wordLower] {
		return "an"
	}

	// check special "a" words first
	if aSpecial[wordLower] {
		return "a"
	}

	// handle uppercase acronyms like FBI, USA
	if anSpecial[wordUpper] {
		return "An"
	}
	if aSpecial[wordUpper] {
		return "A"
	}

	// normal English rule:
	// if word starts with vowel - "an"
	if len(wordLower) > 0 {
		first := wordLower[0]

		// special case: "u-turn", "university", etc.
		if first == 'u' && strings.Contains(wordLower, "-") {
			return "a"
		}

		// vowel check
		if strings.ContainsRune("aeiou", rune(first)) {
			return "an"
		}
	}

	// default fallback
	return "a"
}
