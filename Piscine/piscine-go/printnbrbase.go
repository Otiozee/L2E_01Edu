package piscine

import "github.com/01-edu/z01"

// Function to print the number in the given base with optimized recursion
func PrintNbrBase(nbr int, base string) {
	if !IsValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	baseRunes := []rune(base)
	num := int64(nbr)

	if num < 0 {
		z01.PrintRune('-')
	}
	printBase(uint64(abs(num)), baseRunes)
}

func printBase(num uint64, base []rune) {
	if num == 0 {
		return
	}
	printBase(num/uint64(len(base)), base)
	z01.PrintRune(base[num%uint64(len(base))])
}

func abs(num int64) uint64 {
	if num < 0 {
		return uint64(-num)
	}
	return uint64(num)
}

// Function to check if the base is valid
func IsValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	seen := make(map[rune]bool)
	for _, r := range base {
		if r == '+' || r == '-' || seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}
