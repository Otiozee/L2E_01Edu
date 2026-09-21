package artist

import ("strings"
"github.com/Otiozee/askey/artist/color")

const Reset = "\033[0m"

// GenerateArt converts normal text into ASCII art with user
// preferred color by combining the 8-line banner
// patterns of each character into complete output rows.
func GenerateArt(lines []string,
	bannerMap map[rune][8]string,
	colorId, target string,
) []string {

	var result []string

	colorEverthing := target == ""
	colorId = color.GetColorCode(colorId)

	targets := make(map[rune]bool)
	for _, ch := range target {
		targets[ch] = true
	}

	for i, line := range lines {
		if line == "" {
			if i != 0 {
				result = append(result, "")
			}
			continue
		}

		rows := make([]strings.Builder, 8)

		for idx, ch := range line {
			colorRune := colorEverthing

			if target != "" {
				for start := 0; start <= len(line)-len(target); start++ {
					if line[start:start+len(target)] == target {
						if idx >= start && idx < start+len(target) {
							colorRune = true
							break
						}
					}
				}
			}
			block := bannerMap[ch]

			for r := 0; r < 8; r++ {
				if colorId != "" && colorRune {
					rows[r].WriteString(colorId)
					rows[r].WriteString(block[r])
					rows[r].WriteString(Reset)
				} else {
					rows[r].WriteString(block[r])
				}
			}
		}

		for r := 0; r < 8; r++ {
			result = append(result, rows[r].String())
		}
	}

	return result
}
