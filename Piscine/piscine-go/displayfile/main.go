package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Check the number of arguments
	if len(os.Args) < 2 {
		// If no file argument is provided
		fmt.Println("File name missing")
		return
	} else if len(os.Args) > 2 {
		// If too many arguments are provided
		fmt.Println("Too many arguments")
		return
	}

	// Get the file name from arguments
	fileName := os.Args[1]

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		// If there is an error opening the file
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure the file gets closed after we're done

	// Read and print the content of the file
	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		// If there is an error reading the file
		fmt.Println("Error reading file:", err)
		return
	}
}
