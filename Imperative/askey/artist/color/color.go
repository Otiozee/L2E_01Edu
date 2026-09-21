package color

import (
	"fmt"
	"strconv"
	"strings"
)

// GetColorCode returns the ANSI color code for a named,
// hex, RGB, or supported HSL color.
func GetColorCode(color string) string {
	color = strings.TrimSpace(strings.ToLower(color))
	color = strings.ReplaceAll(color, " ", "")

	// Named colors
	switch color {
	case "black":
		return "\033[30m"
	case "red":
		return "\033[31m"
	case "green":
		return "\033[32m"
	case "yellow":
		return "\033[33m"
	case "purple":
		return "\033[35m"
	case "blue":
		return "\033[34m"
	case "magenta":
		return "\033[35m"
	case "cyan":
		return "\033[36m"
	case "white":
		return "\033[37m"
	case "orange":
		return "\033[38;2;255;165;0m"
	}

	// Hex: #RRGGBB
	if strings.HasPrefix(color, "#") && len(color) == 7 {
		r, err1 := strconv.ParseInt(color[1:3], 16, 64)
		g, err2 := strconv.ParseInt(color[3:5], 16, 64)
		b, err3 := strconv.ParseInt(color[5:7], 16, 64)

		if err1 == nil && err2 == nil && err3 == nil {
			return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
		}
	}

	// RGB: rgb(255,0,0)
	if strings.HasPrefix(color, "rgb(") && strings.HasSuffix(color, ")") {
		content := strings.TrimSuffix(strings.TrimPrefix(color, "rgb("), ")")
		parts := strings.Split(content, ",")

		if len(parts) == 3 {
			r, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
			g, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
			b, err3 := strconv.Atoi(strings.TrimSpace(parts[2]))

			if err1 == nil && err2 == nil && err3 == nil {
				return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
			}
		}
	}

	// hsl(0,0%,0%)
	switch color {
	case "hsl(0,100%,50%)":
		return "\033[38;2;255;0;0m"
	case "hsl(120,100%,50%)":
		return "\033[38;2;0;255;0m"
	case "hsl(240,100%,50%)":
		return "\033[38;2;0;0;255m"
	case "hsl(0,0%,0%)":
		return "\033[30m"
	case "hsl(60,100%,50%)":
		return "\033[33m"
	case "hsl(300,100%,50%)":
		return "\033[35m"
	case "hsl(180,100%,50%)":
		return "\033[36m"
	case "hsl(0,0%,100%)":
		return "\033[37m"
	}

	return ""
}
