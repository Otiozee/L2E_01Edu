package main

import "github.com/01-edu/z01"

func main() {
	// Loop through the ASCII values of 'a' to 'z'
	for r := 'a'; r <= 'z'; r++ {
		// Print each character using z01.PrintRune
		z01.PrintRune(r)
	}
	// Optionally add a newline after printing all characters
	z01.PrintRune('\n')
}
