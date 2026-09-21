package artist

import "strings"

// SplitInput converts literal \n into actual new lines and
// splits the input into separate text lines for processing.
func SplitInput(input string) []string {
	input = strings.ReplaceAll(input, "\\n", "\n")
	return strings.Split(input, "\n")
}
