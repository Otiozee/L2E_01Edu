package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printHelp() {
	helpText := []string{
		"--insert",
		"  -i",
		"\t This flag inserts the string into the string passed as argument.",
		"--order",
		"  -o",
		"\t This flag will behave like a boolean, if it is called it will order the argument.",
	}
	for _, line := range helpText {
		printStr(line)
		z01.PrintRune('\n')
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func orderString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printHelp()
		return
	}

	var insertStr string
	var order bool
	var str string

	for i := 0; i < len(args); i++ {
		arg := args[i]
		if arg == "--help" || arg == "-h" {
			printHelp()
			return
		}

		if startsWith(arg, "--insert=") {
			insertStr = splitOnEqual(arg)[1]
		} else if startsWith(arg, "-i=") {
			insertStr = splitOnEqual(arg)[1]
		} else if arg == "--insert" || arg == "-i" {
			if i+1 < len(args) {
				insertStr = args[i+1]
				i++
			}
		} else if arg == "--order" || arg == "-o" {
			order = true
		} else if !startsWith(arg, "-") {
			str = arg
		}
	}

	if str == "" {
		if insertStr != "" || order {
			if insertStr != "" {
				str = insertStr
				insertStr = ""
			} else {
				printHelp()
				return
			}
		} else {
			printHelp()
			return
		}
	}

	if insertStr != "" {
		str += insertStr
	}

	if order {
		str = orderString(str)
	}

	if str != "" {
		printStr(str)
		z01.PrintRune('\n')
	}
}

func startsWith(s, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if s[i] != prefix[i] {
			return false
		}
	}
	return true
}

func splitOnEqual(s string) []string {
	for i := 0; i < len(s); i++ {
		if s[i] == '=' {
			return []string{s[:i], s[i+1:]}
		}
	}
	return []string{s}
}
