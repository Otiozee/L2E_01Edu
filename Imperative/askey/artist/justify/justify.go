package justify

import (
	"strings"
	"github.com/Otiozee/askey/artist"
)

// JustifyArt generates ASCII art for each word separately and
// distributes the available terminal space evenly between them.
func JustifyArt(
	input string,
	width int,
	bannerMap map[rune][8]string,
	color string,
	target string,
) []string {

	words := strings.Fields(input)

	if len(words) <= 1 {
		return artist.GenerateArt([]string{input}, bannerMap, color, target)
	}

	wordArts := make([][]string, len(words))

	totalWidth := 0

	for i, word := range words {

		wordArts[i] = artist.GenerateArt(
			[]string{word},
			bannerMap,
			color,
			target,
		)
		// for _, row := range wordArts[i] {
		// 	fmt.Println(len(row))
		// }

		// fmt.Println(word, len(wordArts[i][0]))
		totalWidth += len(wordArts[i][0]) // width of this word's ASCII art
	}

	gaps := len(words) - 1

	if width <= totalWidth {
		return artist.GenerateArt(
			[]string{input},
			bannerMap,
			color,
			target,
		)
	}

	space := width - totalWidth
	base := space / gaps
	extra := space % gaps

	result := make([]string, 8)

	for row := 0; row < 8; row++ {

		var b strings.Builder

		for i := 0; i < len(words); i++ {

			b.WriteString(wordArts[i][row])

			if i < gaps {

				n := base
				if extra > 0 {
					n++
					extra--
				}

				b.WriteString(strings.Repeat(" ", n))
			}
		}

		result[row] = b.String()
	}

	return result
}
