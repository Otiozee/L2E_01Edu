package justify

import (
	"os"
	"strings"
	"syscall"
	"unsafe"
)

// AlignArt aligns generated ASCII art according to the user's
// preferred alignment type and the current terminal width.
func AlignArt(rows []string, align string) []string {
	termWidth := TerminalWidth()

	maxWidth := 0
	for _, row := range rows {
		if len(row) > maxWidth {
			maxWidth = len(row)
		}
	}

	if maxWidth >= termWidth {
		return rows
	}

	padding := termWidth - maxWidth

	result := make([]string, len(rows))

	for i, row := range rows {
		switch align {

		case "right":
			result[i] = strings.Repeat(" ", padding) + row

		case "center":
			result[i] = strings.Repeat(" ", padding/2) + row

		default: // left
			result[i] = row
		}
	}

	return result
}

// func maxWidth(rows []string) int {
// 	max := 0
// 	for _, r := range rows {
// 		if len(r) > max {
// 			max = len(r)
// 		}
// 	}
// 	return max
// }

// winsize holds terminal size data from the OS
type winsize struct {
	Row uint16
	Col uint16
}

func TerminalWidth() int {
	ws := &winsize{}

	// syscall.TIOCGWINSZ finds the correct magic number for Linux OR Mac automatically
	_, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		os.Stdout.Fd(),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)

	if errno != 0 {
		return 80
	}
	// fmt.Println(*ws)
	return int(ws.Col)
}
