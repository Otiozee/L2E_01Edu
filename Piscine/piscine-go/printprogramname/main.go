package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	fullpath := os.Args[0]

	lastslash := -1
	for i, c := range fullpath {
		if c == '/' {
			lastslash = i
		}
	}
	for i, c := range fullpath {
		if i > lastslash {
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')
}
