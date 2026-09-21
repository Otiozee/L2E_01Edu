package piscine

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T') // Prints 'T' if the number is negative
	} else {
		z01.PrintRune('F') // Prints 'F' otherwise
	}
	z01.PrintRune('\n') // Prints newline after the output
}
