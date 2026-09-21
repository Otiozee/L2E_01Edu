package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			z01.PrintRune(digits[i/10])
			z01.PrintRune(digits[i%10])
			z01.PrintRune(' ')
			z01.PrintRune(digits[j/10])
			z01.PrintRune(digits[j%10])

			if i != 98 || j != 99 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
