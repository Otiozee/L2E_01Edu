package artist

import (
	"embed"
	"fmt"
	"strings"
)

//go:embed banner/*.txt
var bannerFS embed.FS

// LoadBanner reads the selected banner file (standard, shadow, or thinkertoy) and
// stores each ASCII character pattern in a map for rendering.
func LoadBanner(name string) (map[rune][8]string, error) {
	file := "artist/banner/" + name + ".txt"

	cleanFile := strings.TrimPrefix(file, "artist/")
	data, err := bannerFS.ReadFile(cleanFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read banner style: %v", cleanFile)
	}

	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(content, "\n")

	bannerMap := make(map[rune][8]string)

	Ascii := 32
	for i := 1; i+8 < len(lines); i += 9 {
		var block [8]string
		for j := 0; j < 8; j++ {
			block[j] = lines[i+j]
		}
		bannerMap[rune(Ascii)] = block
		Ascii++
	}
	return bannerMap, nil
}
