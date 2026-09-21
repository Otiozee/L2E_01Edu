package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0') // Xpcial case for 0
		return
	}

	// Create a frequency array to count occurrences of each digit
	frequency := make([]int, 10)
	for n > 0 {
		frequency[n%10]++ // Increase the count for the current digit
		n /= 10           // Remove the last digit from the number
	}

	// Print digits in ascending order based on frequency
	for i := 0; i < 10; i++ {
		for frequency[i] > 0 {
			// Convert the digit to a rune and print it
			z01.PrintRune(rune(i + '0'))
			frequency[i]-- // Decrease the count for the current digit
		}
	}
}
