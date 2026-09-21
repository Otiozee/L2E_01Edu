package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Get command line arguments
	args := os.Args[1:] // Exclude the program name

	// Print arguments in reverse order
	for i := len(args) - 1; i >= 0; i-- {
		for _, ch := range args[i] {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
