## Overview
Go-Reloaded is a sophisticated text processing tool written in Go, designed to parse, modify, and refine input text through a series of modular operations. It tokenizes the input string, applies custom modifiers (such as case transformations, hexadecimal or binary conversions), rebuilds the text, and performs post-processing fixes for punctuation, quotes, and article usage (e.g., "a" vs. "an"). This project demonstrates clean, modular Go code structure, regular expressions for tokenization, and string manipulation techniques.
The tool is ideal for text normalization tasks, educational purposes in learning Go's string handling, or as a base for more advanced natural language processing pipelines. It handles various edge cases like punctuation clustering, quoted phrases, and vowel-based article corrections.

### Features
- Tokenization: Breaks down text into words, punctuation, newlines, and modifier commands enclosed in parentheses.
- Modifiers: Supports commands like (up), (low), (cap) for case changes; (hex) and (bin) for number conversions; with optional counts to apply to previous words.
- Rebuilding: Reassembles tokens into coherent text with proper spacing.
#### Post-Processing:
  - Fixes punctuation spacing.
  - Removes extra spaces inside single quotes.
  - Corrects indefinite articles ("a" to "an" before vowels).

### Pipeline Architecture
```text
Input
  ↓
Tokenize
  ↓
ApplyModifiers
  ↓
Rebuild
  ↓
FixPunctuation
  ↓
FixQuotes
  ↓
FixArticles
  ↓
Output
```

### Program Structure
```text
go-reloaded/
├── helpers/
|   ├── articles.go
|   ├── helpers.go
|   ├── modifiers.go
|   ├── processor.go 
|   ├── punctuation.go 
|   ├── quotes.go
|   ├── rebuild.go
|   └── tokenize.go
├── go.mod
├── main.go
├── processor_test.go
├── README.md
└── sample.txt
```

## Language & Tools

<div align="center">
  <img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" alt="Go Programming Language" width="180"/>
  <br/>
  <strong>Go (Golang)</strong> - Simple, fast, and reliable.
</div>

- Built with the official Go standard library only (no external packages)
- Uses powerful built-in packages: `regexp`, `strings`, `strconv`, `unicode`

## Installation
- Clone the repository: git clone https://acad.learn2earn.ng/git/zeotokpa/go-reloaded.git
- cd go-reloaded

## Testing
- Includes unit tests in processor_test.go to ensure functionality.
- Use go test or go test -v

## Usage
- go run . sample.txt result.txt

## Programmer
Zeotokpa (Zed).