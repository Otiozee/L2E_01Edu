package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		z01.PrintRune('\n')
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]
	seen := make(map[rune]bool)

	for _, c := range str1 {
		if !seen[c] {
			seen[c] = true
			z01.PrintRune(c)
		}
	}
	for _, c := range str2 {
		if !seen[c] {
			seen[c] = true
			z01.PrintRune(c)
		}
	}

	z01.PrintRune('\n')
}
