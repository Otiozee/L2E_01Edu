package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// If no arguments provided, read from stdin
	if len(os.Args) == 1 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			printError(err.Error())
			os.Exit(1)
		}
		return
	}

	// Process file arguments
	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			printError("ERROR: " + err.Error())
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, file)
		file.Close()
		if err != nil {
			printError(err.Error())
			os.Exit(1)
		}
	}
}

func printError(msg string) {
	for _, r := range msg {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
