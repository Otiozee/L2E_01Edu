package main

import (
	"fmt"
	"go-reloaded/helpers"
	"os"
)

// main is the entry point of the program
// It reads an input file, processes the text, and writes to an output file
func main() {

	// check if user provided correct number of arguments
	// expected: program input.txt output.txt
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		os.Exit(1)
	}

	// get input and output file names from command line arguments
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// read the content of the input file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}

	// process the text using helper functions (tokenize, fix, rebuild, etc.)
	result := helpers.ProcessText(string(data))

	// write the final processed result into output file
	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		os.Exit(1)
	}
}
