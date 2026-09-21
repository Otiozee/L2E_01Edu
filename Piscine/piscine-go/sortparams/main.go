package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Function to swap two elements in a slice
func swap(args []string, i, j int) {
	args[i], args[j] = args[j], args[i]
}

// Bubble Sort to sort the arguments in ASCII order
func bubbleSort(args []string) {
	for i := 0; i < len(args)-1; i++ {
		for j := 0; j < len(args)-1-i; j++ {
			if args[j] > args[j+1] {
				swap(args, j, j+1)
			}
		}
	}
}

func main() {
	// Get the command-line arguments (excluding the program name)
	args := os.Args[1:]

	// Sort the arguments in ASCII order
	bubbleSort(args)

	// Print each sorted argument character by character
	for _, arg := range args {
		for _, char := range arg {
			z01.PrintRune(char)
		}
		// Print a newline after each argument
		z01.PrintRune('\n')
	}
}
