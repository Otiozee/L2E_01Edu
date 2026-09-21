package main

import (
	"strings"
	"testing"

	"github.com/Otiozee/askey/artist"
	"github.com/Otiozee/askey/artist/color"
	"github.com/Otiozee/askey/artist/justify"
	"github.com/Otiozee/askey/artist/output"
	"github.com/Otiozee/askey/artist/reverse"
	"github.com/Otiozee/askey/artist/validator"
)

var banner string
var bannerpath = "banner/" + banner + ".txt"

// ---------------------------------------------------------------------------
// Tests for SplitInput
// ---------------------------------------------------------------------------

func TestSplit_EmptyString(t *testing.T) {
	result := artist.SplitInput("")
	if len(result) != 1 || result[0] != "" {
		t.Errorf("SplitInput(\"\") = %v; want [\"\"]", result)
	}
}

func TestSplit_NoNewline(t *testing.T) {
	result := artist.SplitInput("Hello")
	if len(result) != 1 || result[0] != "Hello" {
		t.Errorf("SplitInput(\"Hello\") = %v; want [\"Hello\"]", result)
	}
}

func TestSplit_DoubleNewline(t *testing.T) {
	result := artist.SplitInput(`Hello\n\nThere`)
	expected := []string{"Hello", "", "There"}

	if len(result) != len(expected) {
		t.Fatalf("len=%d; want %d", len(result), len(expected))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("result[%d]=%q; want %q", i, result[i], expected[i])
		}
	}
}

// ---------------------------------------------------------------------------
// Tests for Validator
// ---------------------------------------------------------------------------

func TestValidator_ColorOnly(t *testing.T) {
	Config, err :=
		validator.Validator([]string{
			"prog",
			"--color=red",
			"Hello",
		})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if Config.Color != "red" {
		t.Errorf("color=%q; want red", Config.Color)
	}

	if Config.Target != "" {
		t.Errorf("target=%q; want empty", Config.Target)
	}

	if Config.Input != "Hello" {
		t.Errorf("input=%q; want Hello", Config.Input)
	}

	if Config.Banner != "standard" {
		t.Errorf("banner=%q; want standard", Config.Banner)
	}
}

func TestValidator_ColorTarget(t *testing.T) {
	Config, err :=
		validator.Validator([]string{
			"prog",
			"--color=blue",
			"lo",
			"Hello",
		})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if Config.Color != "blue" {
		t.Errorf("color=%q; want blue", Config.Color)
	}

	if Config.Target != "lo" {
		t.Errorf("target=%q; want lo", Config.Target)
	}

	if Config.Input != "Hello" {
		t.Errorf("input=%q; want Hello", Config.Input)
	}

	if Config.Banner != "standard" {
		t.Errorf("banner=%q; want standard", Config.Banner)
	}
}

func TestValidator_ColorTargetBanner(t *testing.T) {
	Config, err :=
		validator.Validator([]string{
			"prog",
			"--color=green",
			"lo",
			"Hello",
			"shadow",
		})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if Config.Color != "green" {
		t.Errorf("color=%q; want green", Config.Color)
	}

	if Config.Target != "lo" {
		t.Errorf("target=%q; want lo", Config.Target)
	}

	if Config.Input != "Hello" {
		t.Errorf("input=%q; want Hello", Config.Input)
	}

	if Config.Banner != "shadow" {
		t.Errorf("banner=%q; want shadow", Config.Banner)
	}
}

func TestValidator_InvalidColor(t *testing.T) {
	_, err :=
		validator.Validator([]string{
			"prog",
			"--color=pink",
			"Hello",
		})

	if err == nil {
		t.Error("expected error for invalid color")
	}
}

func TestValidator_InvalidFlag(t *testing.T) {
	_, err :=
		validator.Validator([]string{
			"prog",
			"color=red",
			"H",
			"Hello",
		})

	if err == nil {
		t.Error("expected error for invalid color flag")
	}
}

func TestValidator_InvalidBanner(t *testing.T) {
	_, err :=
		validator.Validator([]string{
			"prog",
			"--color=red",
			"H",
			"Hello",
			"matrix",
		})

	if err == nil {
		t.Error("expected error for invalid banner")
	}
}

func TestValidator_InvalidChar(t *testing.T) {
	_, err :=
		validator.Validator([]string{
			"prog",
			"--color=red",
			"Hello\tWorld",
		})

	if err == nil {
		t.Error("expected error for invalid character")
	}
}

// ---------------------------------------------------------------------------
// Tests for GetColorCode
// ---------------------------------------------------------------------------

func TestGetColorCode_Red(t *testing.T) {
	if color.GetColorCode("red") != "\033[31m" {
		t.Error("red color code incorrect")
	}
}

func TestGetColorCode_Green(t *testing.T) {
	if color.GetColorCode("green") != "\033[32m" {
		t.Error("green color code incorrect")
	}
}

func TestGetColorCode_Invalid(t *testing.T) {
	if color.GetColorCode("invalid") != "" {
		t.Error("expected empty string for invalid color")
	}
}

// ---------------------------------------------------------------------------
// Tests for LoadBanner
// ---------------------------------------------------------------------------

func TestLoadBanner_Standard(t *testing.T) {
	bmap, err := artist.LoadBanner("standard")

	if err != nil {
		t.Fatalf("LoadBanner error: %v", err)
	}

	if len(bmap) != 95 {
		t.Errorf("len=%d; want 95", len(bmap))
	}
}

func TestLoadBanner_Shadow(t *testing.T) {
	_, err := artist.LoadBanner("shadow")

	if err != nil {
		t.Fatalf("LoadBanner(shadow) error: %v", err)
	}
}

func TestLoadBanner_Thinkertoy(t *testing.T) {
	_, err := artist.LoadBanner("thinkertoy")

	if err != nil {
		t.Fatalf("LoadBanner(thinkertoy) error: %v", err)
	}
}

func TestLoadBanner_Invalid(t *testing.T) {
	_, err := artist.LoadBanner("fakebanner")

	if err == nil {
		t.Error("expected error")
	}
}

// ---------------------------------------------------------------------------
// Tests for GenerateArt
// ---------------------------------------------------------------------------

func TestGenerateArt_EmptyLine(t *testing.T) {
	bmap, _ := artist.LoadBanner("standard")

	art := artist.GenerateArt([]string{""}, bmap, "", "")

	if art == nil {
		return
	}

	for _, row := range art {
		if row != "" {
			t.Errorf("expected empty rows, got %v", art)
		}
	}
}

func TestGenerateArt_A(t *testing.T) {
	bmap, _ := artist.LoadBanner("standard")

	art := artist.GenerateArt([]string{"A"}, bmap, "", "")

	if len(art) != 8 {
		t.Fatalf("len=%d; want 8", len(art))
	}
}

func TestGenerateArt_ColorAll(t *testing.T) {
	bmap, _ := artist.LoadBanner("standard")

	art := artist.GenerateArt([]string{"A"}, bmap, "red", "")

	found := false

	for _, row := range art {
		if strings.Contains(row, "\033[31m") {
			found = true
			break
		}
	}

	if !found {
		t.Error("expected red ANSI code")
	}
}

func TestGenerateArt_TargetColor(t *testing.T) {
	bmap, _ := artist.LoadBanner("standard")

	art := artist.GenerateArt([]string{"AB"}, bmap, "green", "A")

	found := false

	for _, row := range art {
		if strings.Contains(row, "\033[32m") {
			found = true
			break
		}
	}

	if !found {
		t.Error("expected green ANSI code")
	}
}

func TestGenerateArt_NoColor(t *testing.T) {
	bmap, _ := artist.LoadBanner("standard")

	art := artist.GenerateArt([]string{"A"}, bmap, "", "")

	for _, row := range art {
		if strings.Contains(row, "\033[") {
			t.Error("unexpected ANSI code found")
		}
	}
}

// ---------------------------------------------------------------------------
// Integration Tests
// ---------------------------------------------------------------------------

func TestRender_EmptyInput(t *testing.T) {
	bmap, _ := artist.LoadBanner("standard")

	lines := artist.SplitInput("")
	art := artist.GenerateArt(lines, bmap, "", "")

	artist.Render(art)
}

func TestRender_Hello(t *testing.T) {
	bmap, _ := artist.LoadBanner("standard")

	lines := artist.SplitInput("Hello")
	art := artist.GenerateArt(lines, bmap, "red", "")

	artist.Render(art)
}

// ---------------------------------------------------------------------------
// Tests for Output
// ---------------------------------------------------------------------------

func TestOutputArt_Hello(t *testing.T) {
	bmap, _ := artist.LoadBanner("standard")
	lines := artist.SplitInput("Hello")
	art := artist.GenerateArt(lines, bmap, "", "")
	err := output.OutputArt(art, "a.txt")
	if err != nil {
		t.Error("failed to create output")
	}
}

func TestOutputArt_EmptyInput(t *testing.T) {
	bmap, _ := artist.LoadBanner("standard")
	lines := artist.SplitInput("")
	art := artist.GenerateArt(lines, bmap, "", "")
	err := output.OutputArt(art, "b.txt")
	if err != nil {
		t.Error("failed to create output")
	}
}

// ---------------------------------------------------------------------------
// Tests for Reverse
// ---------------------------------------------------------------------------

func TestReverseArt_A(t *testing.T) {
	bmap, _ := artist.LoadBanner("standard")

	art := artist.GenerateArt([]string{"A"}, bmap, "", "")

	txt, err := reverse.ReverseArt(art, bmap)

	if err != nil {
		t.Fatal(err)
	}

	if txt != "A" {
		t.Errorf("got %q want A", txt)
	}
}

func TestReverseArt_Hello(t *testing.T) {
	bmap, _ := artist.LoadBanner("standard")

	art := artist.GenerateArt([]string{"Hello"}, bmap, "", "")

	txt, err := reverse.ReverseArt(art, bmap)

	if err != nil {
		t.Fatal(err)
	}

	if txt != "Hello" {
		t.Errorf("got %q want Hello", txt)
	}
}

// ---------------------------------------------------------------------------
// Tests for Alignment
// ---------------------------------------------------------------------------

func TestAlign_Left(t *testing.T) {
	art := []string{"Hello"}

	out := justify.AlignArt(art, "left")

	if out[0] != "Hello" {
		t.Error("left alignment failed")
	}
}

func TestAlign_Center(t *testing.T) {
	art := []string{"Hello"}

	out := justify.AlignArt(art, "center")

	if len(out) != 1 {
		t.Error("unexpected output")
	}
}

func TestAlign_Right(t *testing.T) {
	art := []string{"Hello"}

	out := justify.AlignArt(art, "right")

	if len(out) != 1 {
		t.Error("unexpected output")
	}
}

func TestJustifyArt_OneWord(t *testing.T) {
	bmap, _ := artist.LoadBanner("standard")

	art := justify.JustifyArt(
		"Hello",
		100,
		bmap,
		"",
		"",
	)

	if len(art) != 8 {
		t.Fatalf("got %d rows; want 8", len(art))
	}
}