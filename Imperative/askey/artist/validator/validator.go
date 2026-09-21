package validator

import (
	"fmt"
	"strings"
	"github.com/Otiozee/askey/artist/color"
)

type Config struct {
	Color   string
	Target  string
	Input   string
	Banner  string
	Output  string
	Align   string
	Reverse string
}

// Validator verifies if: the correct number of arguments is provided,
// or the banner name is valid
// or if the flag is valid
// or if color is valid and,
// or the input contains only supported printable ASCII characters.
func Validator(args []string) (Config, error) {
	cfg := Config{
		Banner: "standard",
	}

	var positional []string

	for _, arg := range args[1:] {
		switch {
		case strings.HasPrefix(arg, "--color="):
			cfg.Color = strings.TrimPrefix(arg, "--color=")

		case strings.HasPrefix(arg, "--output="):
			cfg.Output = strings.TrimPrefix(arg, "--output=")

		case strings.HasPrefix(arg, "--align="):
			cfg.Align = strings.TrimPrefix(arg, "--align=")

		case strings.HasPrefix(arg, "--reverse="):
			cfg.Reverse = strings.TrimPrefix(arg, "--reverse=")

		default:
			positional = append(positional, arg)
		}
	}

	// align mode
	if cfg.Align == "" {
		cfg.Align = "left"
	}

	switch cfg.Align {
	case "left", "right", "center", "justify":
	default:
		return Config{}, fmt.Errorf("invalid align type: %s\n"+
			"valid type:\n"+
			"left | right | center | justify", cfg.Align)
	}

	// reverse mode
	if cfg.Reverse != "" {
		return cfg, nil
	}

	switch len(positional) {
	case 1:
		cfg.Input = positional[0]

	case 2:
		if isBanner(positional[1]) {
			cfg.Input = positional[0]
			cfg.Banner = positional[1]
		} else {
			cfg.Target = positional[0]
			cfg.Input = positional[1]
		}

	case 3:
		cfg.Target = positional[0]
		cfg.Input = positional[1]
		cfg.Banner = positional[2]

	default:
		return Config{}, fmt.Errorf(
			"invalid number of arguments\n" +
				"Usage:\n" +
				"go run . [string]\n" +
				"go run . [string] [banner]\n" +
				"go run . [flag] [string] [banner]",
		)
	}

	if !isBanner(cfg.Banner) {
		return Config{}, fmt.Errorf(
			"invalid banner '%s'\n"+
				"available banners: standard | shadow | thinkertoy",
			cfg.Banner,
		)
	}

	if cfg.Color != "" && color.GetColorCode(cfg.Color) == "" {
		return Config{}, fmt.Errorf(
			"invalid color '%s'\n"+
				"availabe colors:\n"+
				"red | blue | green | white | black | purple | margenta | yellow | cyan | orange",
			cfg.Color,
		)
	}

	for _, r := range cfg.Input {
		if r == '\n' {
			continue
		}

		if r < 32 || r > 126 {
			return Config{}, fmt.Errorf(
				"invalid ASCII character %q (byte '%d')",
				r, r,
			)
		}
	}

	return cfg, nil
}

func isBanner(banner string) bool {
	switch banner {
	case "standard", "shadow", "thinkertoy":
		return true
	}
	return false
}
