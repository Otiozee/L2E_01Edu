package artist

import "fmt"

// Render prints the generated ASCII-art lines to the terminal.
func Render(art []string) {
	for _, line := range art {
		fmt.Println(line)
	}
}
