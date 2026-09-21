> This project is designed as both a standalone CLI tool and a reusable Go package.

# Overview
ASkey is a feature-rich, well-structured project that goes beyond the standard requirements and demonstrates good handling of file writing, validation, for string-to-art conversion. It is designed to read user input string and print to the terminal or write out the the art version of the input into a user preferred .txt file using available ascii banner style. It loads banner file, reads user input, validates both the loaded banner and the user input, generate the intended art with user preferred color and target character, and renders it accordingly.

This project demonstrates clean Go code structure, `os` and `embed` for file compilation and handling, and `string` for manipulation techniques. The tool is ideal for text-to-art conversion tasks, educational purposes in learning Go's string handling, structs, mapping,  as well as slicing. It handles various edge cases like `trailing newlines`, `leading newline`, `newlines in between strings`, and `invalid ascii characters`, `whole text coloring`, `selected or targeted coloring`.

---

## Features
- `main:` this is the executor of the entire application. It does that through the ProcessArt().
```text
main()
 └── ProcessArt()
        ├── Validator()
        ├── LoadBanner()
        ├── ReverseArt() (optional)
        ├── SplitInput()
        ├── JustifyArt() / GenerateArt()
        │      └── GetColorCode()
        ├── AlignArt() (optional)
        ├── OutputArt() (optional)
        └── Render()
```
- `Validator:` Checks if:
    * the correct number of arguments is provided,
    * the banner name is valid 
    * the alignment option is valid,
    * the output and reverse options are correctly supplied,
    * the color is valid or available color, and,
    * the input contains only supported printable ASCII characters.

- `LoadBanner:` Reads the selected banner file (standard, shadow, or thinkertoy) and stores each ASCII character pattern in a map for rendering.

- `ReverseArt`: Converts ASCII art stored in a file back into normal text.

- `SplitInput:` Converts literal \n into actual new lines and splits the input into separate text lines for processing.

- `GenerateArt:` Converts normal text into ASCII art by combining the 8-line banner patterns of each character and color preference into complete output rows.

- `GetColorCode:` Contains all available colors for the project
and returns the corresponding ANSI color code according to the user preferences.

- `JustifyArt`: Produces justified ASCII art according to the current terminal width.

- `AlignArt`: Applies left, right, center, or justify alignment to the generated ASCII art.

- `OutputArt`: Writes the generated ASCII art to a file when the --output option is used.

- `Render:` Prints the generated ASCII-art lines to the terminal.

- Reusable Go package (`artist`)

---

## Pipeline Architecture
```text
User Input (CLI Arguments)
        │
        ▼
main()
        │
        ▼
ProcessArt()
        │
        ├── Validator()
        │      ├── validates argument count
        │      ├── validates banner name
        │      ├── validates alignment
        │      ├── validates output/reverse options
        │      ├── validates color
        │      └── validates ASCII characters
        │
        ▼
LoadBanner()
        │
        ├── reads banner file
        ├── extracts ASCII blocks
        └── stores them in a map
        │
        ▼
Reverse Mode?
        │
   Yes ─┴──────────────► ReverseArt()
        │                     │
   No   ▼                     └── prints decoded text
SplitInput()
        │
        ├── converts "\n" → newline
        └── splits text into lines
        │
        ▼
Alignment = justify?
        │
   Yes ─┴──────────────► JustifyArt()
        │                     │
   No   ▼                     └── calls GetColorCode()
GenerateArt()
        │
        ├── loops through text
        ├── retrieves ASCII blocks
        ├── calls GetColorCode()
        └── builds colored ASCII art
        │
        ▼
AlignArt()
        │
        ├── left
        ├── right
        ├── center
        └── justify
        │
        ▼
OutputArt()
        │
        ├── writes to file (optional)
        └── returns generated art
        │
        ▼
Render()
        │
        └── prints ASCII art to terminal
```

---

## Program Structure
```text
askey
│
├── artist/
│   ├── banner/
│   │   ├── shadow.txt
│   │   ├── standard.txt
│   │   └── thinkertoy.txt
│   │
│   ├── color/
│   │   └── color.go
│   │
│   ├── justify/
│   │   ├── align.go
│   │   └── justify.go
│   │
│   ├── output/
│   │   └── outputart.go
│   │
│   ├── processor/
│   │   └── processor.go
│   │
│   ├── reverse/
│   │   └── reverse.go
│   │
│   ├── validator/
│   │   └── validator.go
│   │
│   ├── artgenerator.go
│   ├── loader.go
│   ├── renderer.go
│   └── split.go
│
├── art_test.go
├── go.mod
├── main.go
├── README.md
└── shell_test.sh
```
---

## Language & Tools

<div align="center">
  <img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" alt="Go Programming Language" width="180"/>
  <br/>
  <strong>Go (Golang)</strong> - Simple, fast, and reliable.
</div>

- Built with the official Go standard library only (no external packages)
- Uses powerful built-in packages: `os`, `strings`, `fmt`, `strconv`, `log`, `embed`

---

## Installation
- Clone the repository:
```bash
git clone https://github.com/Otiozee/askey.git
```
- cd ascii-art-engine

---

## Testing
- Includes unit tests to ensure functionality.
- Run test:
```bash
go test -v
```
---

## Usage

### Default Banner

```bash
go run . "Hello"
```

### Custom Banner

```bash
go run . "Hello" shadow
```

### Colored
```bash
go run . --color=red "Hello"
```

### Colored with target text

```bash
go run . --color=green Hello "Hello\nWorld"
```

### Center Alignment

```bash
go run . --align=center "Hello"
```

### Right Alignment

```bash
go run . --align=right "Hello"
```

### Left Alignment

```bash
go run . --align=left "Hello"
```

### Justify Alignment
```bash
go run . --align=justify "Hello World"
```
### Save output file
```bash
go run . --output=result.txt "Hello"
```
### Reverse Art
```bash
go run . --reverse=result.txt
```

### Run muiltiple examples from shell file
```bash
chmod +x shell_test.sh
```
```bash
./shell_test.sh
```
---

## Using as a Package

This project can be imported and reused in other Go applications.

### Public API

The following exported functions can be imported and reused in other Go projects.

```go
processor.ProcessArt()

artist.LoadBanner(...)
artist.SplitInput(...)
artist.GenerateArt(...)
artist.Render(...)

validator.Validator(...)

color.GetColorCode(...)

justify.AlignArt(...)
justify.JustifyArt(...)
justify.TerminalWidth(...)

reverse.ReverseArt(...)

output.OutputArt(...)
```

Import the packages you need:

```go
import (
	"github.com/Otiozee/askey/artist"
	"github.com/Otiozee/askey/artist/processor"
)
```

Example:

```go
package main

import (
	"log"

	artist "github.com/Otiozee/askey/artist"
	"github.com/Otiozee/askey/artist/processor"
)

func main() {
	// Execute the application
	processor.ProcessArt()

	// Or reuse individual components
	bannerMap, err := artist.LoadBanner("standard")
	if err != nil {
		log.Fatal(err)
	}

	lines := artist.SplitInput("Hello")
	art := artist.GenerateArt(lines, bannerMap, "", "")
	artist.Render(art)
}
```
---

## Program Team
- [Zed](https://github.com/otiozee)