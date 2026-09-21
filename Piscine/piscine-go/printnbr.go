package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		if n == -n {
			z01.PrintRune('-')
			min := "9223372036854775808"
			for _, r := range min {
				z01.PrintRune(r)
			}
			return
		}
		z01.PrintRune('-')
		n = -n
	}
	if n >= 10 {
		PrintNbr(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}
