package piscine

import "github.com/01-edu/z01"

func DealAPackOfCards(deck []int) {
	player := 1
	i := 0
	for player <= 4 {
		z01.PrintRune('P')
		z01.PrintRune('l')
		z01.PrintRune('a')
		z01.PrintRune('y')
		z01.PrintRune('e')
		z01.PrintRune('r')
		z01.PrintRune(' ')
		if player == 1 {
			z01.PrintRune('1')
		} else if player == 2 {
			z01.PrintRune('2')
		} else if player == 3 {
			z01.PrintRune('3')
		} else if player == 4 {
			z01.PrintRune('4')
		}
		z01.PrintRune(':')
		z01.PrintRune(' ')

		count := 0
		for count < 3 {
			n := deck[i]
			if n == 10 {
				z01.PrintRune('1')
				z01.PrintRune('0')
			} else if n == 11 {
				z01.PrintRune('1')
				z01.PrintRune('1')
			} else if n == 12 {
				z01.PrintRune('1')
				z01.PrintRune('2')
			} else if n == 9 {
				z01.PrintRune('9')
			} else if n == 8 {
				z01.PrintRune('8')
			} else if n == 7 {
				z01.PrintRune('7')
			} else if n == 6 {
				z01.PrintRune('6')
			} else if n == 5 {
				z01.PrintRune('5')
			} else if n == 4 {
				z01.PrintRune('4')
			} else if n == 3 {
				z01.PrintRune('3')
			} else if n == 2 {
				z01.PrintRune('2')
			} else if n == 1 {
				z01.PrintRune('1')
			}

			if count < 2 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			count++
			i++
		}

		z01.PrintRune('\n')
		player++
	}
}
