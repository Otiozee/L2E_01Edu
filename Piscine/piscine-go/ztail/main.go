package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 3 || args[0] != "-c" {
		return
	}

	n := 0
	for _, ch := range args[1] {
		n = n*10 + int(ch-'0')
	}

	files := args[2:]
	errorOccurred := false
	multiple := len(files) > 1

	for i, fileName := range files {
		content, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println(err) // Correct error message
			errorOccurred = true
			continue
		}

		if multiple {
			if i > 0 {
				fmt.Println()
			}
			fmt.Printf("==> %s <==\n", fileName)
		}

		if n <= len(content) {
			fmt.Print(string(content[len(content)-n:]))
		} else {
			fmt.Print(string(content))
		}
	}

	if errorOccurred {
		os.Exit(1)
	}
}
