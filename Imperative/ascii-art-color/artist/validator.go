package artist

import (
	"fmt"
	"strings"
)

// Validator verifies if: the correct number of arguments is provided,
// or the banner name is valid
// or if color is valid and,
// or the input contains only supported printable ASCII characters.
func Validator(Args []string) (string, string, string, string, error) {
	if len(Args) < 3 || len(Args) > 5 {
		return "", "", "", "", fmt.Errorf("Invalid number of arguments!\n" +
			"Usage: go run . --color=[color] [letters to color] [string] [banner]\n" +
			"Example: go run . --color=red H \"Hello\" shadow")
	}

	color := Args[1]
	target := ""
	input := Args[2]
	banner := "standard"

	if len(Args) == 4 {
		color = Args[1]
		target = Args[2]
		input = Args[3]
		banner = "standard"
	}

	if len(Args) == 5 {
		color = Args[1]
		target = Args[2]
		input = Args[3]
		banner = Args[4]
	}

	valid := map[string]bool{
		"standard":   true,
		"thinkertoy": true,
		"shadow":     true,
	}

	if !valid[banner] {
		return "", "", "", "", fmt.Errorf("Invalid banner '%v' entered.\n"+
			"Available banners: standard | shadow | thinkertoy", banner)
	}

	color = strings.TrimPrefix(color, "--color=")

	validColor := GetColorCode(color)

	if validColor == "" {
		return "", "", "", "", fmt.Errorf("Invalid color '%v' selected.\n"+
			"Valid color names:\n"+
			"red | blue | green | white | black | purple | margenta | yellow | cyan | orange",
			color)
	}
	for _, i := range input {
		if i == '\n' {
			continue
		}
		if i < 32 || i > 126 {
			return "", "", "", "", fmt.Errorf(
				"invalid ASCII character %q (byte %d)",
				i,
				i,
			)
		}
	}
	return color, target, input, banner, nil
}
