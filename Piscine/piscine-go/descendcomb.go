package piscine

import (
	"github.com/01-edu/z01"
)

func DescendComb() {
	first := true
	for i := 99; i >= 1; i-- {
		for j := i - 1; j >= 0; j-- {
			if !first {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			printTwoDigit(i)
			z01.PrintRune(' ')
			printTwoDigit(j)
			first = false
		}
	}
}

func printTwoDigit(n int) {
	z01.PrintRune(rune(n/10) + '0')
	z01.PrintRune(rune(n%10) + '0')
}
