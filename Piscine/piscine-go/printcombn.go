package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n < 1 || n > 9 {
		z01.PrintRune('E') // Error indication
		z01.PrintRune('\n')
		return
	}
	comb := make([]int, n)
	for i := 0; i < n; i++ {
		comb[i] = i
	}
	for {
		printCombination(comb, n)
		if !nextCombination(comb, n) {
			break
		}
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
	z01.PrintRune('\n')
}

func printCombination(comb []int, n int) {
	for i := 0; i < n; i++ {
		z01.PrintRune(rune(comb[i] + '0'))
	}
}

func nextCombination(comb []int, n int) bool {
	for i := n - 1; i >= 0; i-- {
		if comb[i] < 10-n+i {
			comb[i]++
			for j := i + 1; j < n; j++ {
				comb[j] = comb[j-1] + 1
			}
			return true
		}
	}
	return false
}
