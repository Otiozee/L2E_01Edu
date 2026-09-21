package main

import (
	"go-reloaded/helpers"
	"testing"
)

// TestHex checks if hexadecimal numbers are correctly converted to decimal
func TestHex(t *testing.T) {
	input := "1E (hex) files were added"
	expected := "30 files were added"

	if helpers.ProcessText(input) != expected {
		t.Fatal("Hex conversion failed")
	}
}

// TestBin checks if binary numbers are correctly converted to decimal
func TestBin(t *testing.T) {
	input := "It has been 10 (bin) years"
	expected := "It has been 2 years"

	if helpers.ProcessText(input) != expected {
		t.Fatal("Binary conversion failed")
	}
}

// TestUpMulti checks if multiple words are converted to uppercase
func TestUpMulti(t *testing.T) {
	input := "This is so exciting (up, 2)"
	expected := "This is SO EXCITING"

	if helpers.ProcessText(input) != expected {
		t.Fatal("Up multi failed")
	}
}

// TestQuotes checks if extra spaces inside quotes are removed
func TestQuotes(t *testing.T) {
	input := "I am exactly how they describe me: ' awesome '"
	expected := "I am exactly how they describe me: 'awesome'"

	if helpers.ProcessText(input) != expected {
		t.Fatal("Quote spacing failed")
	}
}

// TestArticle checks if "a" is corrected to "an" when needed
func TestArticle(t *testing.T) {
	input := "There it was. A amazing rock!"
	expected := "There it was. An amazing rock!"

	if helpers.ProcessText(input) != expected {
		t.Fatal("Article fix failed")
	}
}

// TestPunctuation checks if spacing before punctuation is fixed
func TestPunctuation(t *testing.T) {
	input := "I was thinking ... You were right"
	expected := "I was thinking... You were right"

	if helpers.ProcessText(input) != expected {
		t.Fatal("Punctuation failed")
	}
}
