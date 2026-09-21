package processor

import (
	"fmt"
	"log"
	"os"
	artist "github.com/Otiozee/askey/artist"
	"github.com/Otiozee/askey/artist/validator"
	"github.com/Otiozee/askey/artist/justify"
	"github.com/Otiozee/askey/artist/reverse"
	"github.com/Otiozee/askey/artist/output"
)

// ProcessArt is the controller of the entire application.
// It follows a definite workflow that combines all the functions.
func ProcessArt() []string {
	Config, err := validator.Validator(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	bannerMap, err := artist.LoadBanner(Config.Banner)
	if err != nil {
		log.Fatal(err)
	}

	if Config.Reverse != "" {
		content, err := os.ReadFile(Config.Reverse)
		if err != nil {
			log.Fatalf("failed to read reverse file: %v", err)
		}

		artRows := artist.SplitInput(string(content))

		for len(artRows) > 0 && artRows[len(artRows)-1] == "" {
			artRows = artRows[:len(artRows)-1]
		}

		plainText, err := reverse.ReverseArt(artRows, bannerMap)
		if err != nil {
			log.Fatalf("reverse failed: %v", err)
		}

		fmt.Println(plainText)
		return nil
	}

	var art []string

	if Config.Align == "justify" {
		art = justify.JustifyArt(
			Config.Input,
			justify.TerminalWidth(),
			bannerMap,
			Config.Color,
			Config.Target,
		)
	} else {
		lines := artist.SplitInput(Config.Input)

		art = artist.GenerateArt(
			lines,
			bannerMap,
			Config.Color,
			Config.Target,
		)

		art = justify.AlignArt(art, Config.Align)
	}

	if err := output.OutputArt(art, Config.Output); err != nil {
		log.Fatal(err)
	}

	return art
}
