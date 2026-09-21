package output

import (
	"fmt"
	"os"
	"strings"
)

// OutputArt writes out the generated art into a file.
func OutputArt(art []string, filename string) error {
	content := strings.Join(art, "\n")

	if filename == "" {
		fmt.Println(content)
		return nil
	}

	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("error writing '%s' %w file", filename, err)
	}
	return nil
}
