package reverse

import "errors"

// ReverseArt() scans the ASCII-art from left to right, repeatedly
// comparing the current 8-row block against every character in the banner
// map until it finds a match, then rebuilds the original text character by character.
func ReverseArt(art []string, bmap map[rune][8]string) (string, error) {
	var res string
	for i := 0; i < len(art); i += 8 {
		if art[i] == "" {
			res += "\n"
			continue
		}
		offset := 0
		for offset < len(art[i]) {
			found := false

			for r, block := range bmap {
				if r == ' ' {
					continue
				} // Skip space here to prioritize real letter outlines
				w := len(block[0])
				if offset+w <= len(art[i]) && match(art, i, offset, block) {
					res += string(r)
					offset += w
					found = true
					break
				}
			}

			if !found {
				if block, exists := bmap[' ']; exists {
					w := len(block[0])
					if offset+w <= len(art[i]) && match(art, i, offset, block) {
						res += " "
						offset += w
						found = true
					}
				}
			}

			if !found {
				if art[i][offset] == ' ' {
					res += " "
					offset++
					continue
				}
				return "", errors.New("err")
			}
		}
		if i+8 < len(art) {
			res += `\n`
			continue
		}
		break
	}
	return res, nil
}

// func match() compares the current piece of the ASCII art against
// a candidate character and returns either: true of false so that
// ReverseArt() can figure out which letter is sitting at the current offset.
func match(art []string, s, o int, b [8]string) bool {
	for r := 0; r < 8; r++ {
		if s+r >= len(art) || o+len(b[r]) > len(art[s+r]) || art[s+r][o:o+len(b[r])] != b[r] {
			return false
		}
	}
	return true
}
