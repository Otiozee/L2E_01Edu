package artist

import (
	"log"
	"os"
)

// ProcessArt is the controller of the entire application.
// It follows a definite workflow that combines all the functions.
func ProcessArt() []string {
	color, target, input, banner, err := Validator(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	bannerMap, err := LoadBanner(banner)
	if err != nil {
		log.Fatal(err)
	}

	lines := SplitInput(input)

	art := GenerateArt(lines, bannerMap, color, target)

	Render(art)

	return art
}
