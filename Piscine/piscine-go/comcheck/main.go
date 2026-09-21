package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	alerts := []string{"01", "galaxy", "galaxy 01"}

	for _, arg := range args {
		for _, alert := range alerts {
			if arg == alert {
				os.Stdout.WriteString("Alert!!!\n")
				return
			}
		}
	}
}
