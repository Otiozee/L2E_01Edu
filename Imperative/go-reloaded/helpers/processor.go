package helpers

// ProcessText is the main pipeline of the program.
// It takes raw input text and applies all transformations step by step.
func ProcessText(input string) string {

	// Step 1: split input text into words/tokens
	tokens := Tokenize(input)

	// Step 2: apply special modifiers like (hex), (bin), (up), etc.
	tokens = ApplyModifiers(tokens)

	// Step 3: rebuild tokens back into a normal sentence string
	text := Rebuild(tokens)

	// Step 4: fix punctuation spacing (like "hello ?" to "hello?")
	text = FixPunctuation(text)

	// Step 5: fix spacing inside quotes (like " hello " to "hello")
	text = FixQuotes(text)

	// Step 6: fix incorrect use of "a" and "an"
	text = FixArticles(text)

	// return final processed text
	return text
}
