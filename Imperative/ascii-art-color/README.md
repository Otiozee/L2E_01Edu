> This project is designed as both a standalone CLI tool and a reusable Go package.

# Overview
ASCII-ART-COLOR is a feature-rich, well-structured project that goes beyond the standard requirements and demonstrates good handling of colors, validation, for string-to-art conversion, designed to read user input string and print out the the art version of the input using available ascii banner files. It loads banner file, reads user input, validates both the loaded banner and the user input, generate the intended art with user preferred color and target character, and renders it accordingly.

This project demonstrates clean Go code structure, `os` and `embed` for file compilation and handling, and `string` for manipulation techniques. The tool is ideal for text-to-art conversion tasks, educational purposes in learning Go's string handling, mapping,  as well as slicing. It handles various edge cases like `trailing newlines`, `leading newline`, `newlines in between strings`, and `invalid ascii characters`, `whole text coloring`, `selected or targeted coloring`.

---

## Features
- `main:` this is the executor of the entire application. It does that through the ProcessArt().
```text
main()
 └── ProcessArt()
        ├── Validator()
        ├── LoadBanner()
        ├── SplitInput()
        ├── GenerateArt()
        |     └── GetColorCode()
        └── Render()
```
- `Validator:` Checks if:
    * the correct number of arguments is provided,
    * the banner name is valid 
    * the color is valid or available color, and,
    * the input contains only supported printable ASCII characters.

- `LoadBanner:` Reads the selected banner file (standard, shadow, or thinkertoy) and stores each ASCII character pattern in a map for rendering.

- `SplitInput:` Converts literal \n into actual new lines and splits the input into separate text lines for processing.

- `GenerateArt:` Converts normal text into ASCII art by combining the 8-line banner patterns of each character and color preference into complete output rows.

- `GetColorCode:` Contains all available colors for the project
and returns the corresponding ANSI color code according to the user preferences.

- `Render:` Prints the generated ASCII-art lines to the terminal.

- Reusable Go package (`artist`)

---

## Pipeline Architecture
```text
User Input (from CLI Arguments)
        │
        |
main()
        |
        ├── executor
        |
        |
ProcessArt()
        │
        └── controller
        │
        |
Validator()
        │
        ├── validates argument count
        ├── validates banner name
        ├── validates color name and ANSI color code
        └── validates ASCII characters
        │
        |
LoadBanner()
        │
        ├── reads banner file
        ├── extracts ASCII blocks
        └── stores them in a map
        │
        |
SplitInput()
        │
        ├── converts "\n" -> actual newline
        └── splits text into lines
        │
        |
GenerateArt()
        │
        ├── loops through each line
        ├── gets ASCII block per character
        ├── calls GetColorCode to generate color
        │        └── returns the ANSI color code
        ├── merges rows horizontally
        └── creates final ASCII-art-color output
Render()
        │
        └── prints generated art to terminal
```

---

## Program Structure
```text
ascii-art-color
│
├── artist/
|    ├── banner/
|    |    ├── shadow.txt
|    |    ├── standard.txt
|    |    └── thinkertoy.txt
|    ├── artgenerator.go
|    ├── color.go
|    ├── loader.go
|    ├── processor.go
|    ├── renderer.go
|    ├── split.go
|    └── validator.go
├── art_test.go
├── go.mod
├── main.go
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
git clone https://acad.learn2earn.ng/git/zeotokpa/ascii-art-color.git
```
- cd ascii-art

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
go run . --color=red "Hello"
```

### Custom Banner

```bash
go run . --color=blue "Hello" shadow
```

### Multiline Input with target

```bash
go run . --color=green Hello "Hello\nWorld"
```

## Examples

### Standard

```bash
go run . --color=cyan "Hello" standard
```

### Shadow

```bash
go run . --color=purple "Hello" shadow
```

### Thinkertoy

```bash
go run . --color=yellow "Hello" thinkertoy
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
These functions can be imported into other Go projects.

```go
artist.ProcessArt()
artist.GenerateArt(...)
artist.LoadBanner(...)
artist.GetColorCode(...)
artist.SplitInput(...)
artist.Render(...)
artist.Validator(...)
```

```go
import "github.com/otiozee/ascii-art-color/artist"
```

Example:

```go
package main

import "github.com/otiozee/ascii-art-color/artist"

func main() {
	artist.ProcessArt()
}
```
---

## Program Team
- zeotokpa