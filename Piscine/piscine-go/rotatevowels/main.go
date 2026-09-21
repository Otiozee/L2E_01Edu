package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		z01.PrintRune('\n')
		return
	}

	// Collect all vowels and their positions across all arguments
	var allVowels []rune
	var vowelPositions []struct {
		argIndex int
		pos      int
	}

	for argIndex, arg := range args {
		for pos, r := range arg {
			if isVowel(r) {
				allVowels = append(allVowels, r)
				vowelPositions = append(vowelPositions, struct {
					argIndex int
					pos      int
				}{argIndex, pos})
			}
		}
	}

	// Mirror the vowels
	for i := 0; i < len(allVowels)/2; i++ {
		j := len(allVowels) - 1 - i
		allVowels[i], allVowels[j] = allVowels[j], allVowels[i]
	}

	// Reconstruct the arguments with mirrored vowels
	outputArgs := make([]string, len(args))
	copy(outputArgs, args)

	for i, pos := range vowelPositions {
		argRunes := []rune(outputArgs[pos.argIndex])
		argRunes[pos.pos] = allVowels[i]
		outputArgs[pos.argIndex] = string(argRunes)
	}

	// Print the result
	for i, arg := range outputArgs {
		for _, r := range arg {
			z01.PrintRune(r)
		}
		if i < len(outputArgs)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
