package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	upper := false

	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	printed := false

	for _, arg := range args {
		n := 0
		valid := true

		for _, ch := range arg {
			if ch >= '0' && ch <= '9' {
				n = n*10 + int(ch-'0')
			} else {
				valid = false
				break
			}
		}
		if valid && n >= 1 && n <= 26 {
			var letter rune
			if upper {
				letter = rune('A' + n - 1)
			} else {
				letter = rune('a' + n - 1)
			}
			z01.PrintRune(letter)
			printed = true
		} else {
			z01.PrintRune(' ')
			printed = true
		}
	}
	if printed {
		z01.PrintRune('\n')
	}
}
