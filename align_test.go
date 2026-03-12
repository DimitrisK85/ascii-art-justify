package main

import (
	"ascii-art-justify/internal/banner"
	"ascii-art-justify/internal/converter"
	"reflect"
	"strings"
	"testing"
)

func TestAlignLeft(t *testing.T) {
	input := []string{
		"* ",
		"**",
	}
	got := alignLeft(input, 4)
	want := []string{
		"*   ",
		"**  ",
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("alignLeft line %d = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestAlignRight(t *testing.T) {
	input := []string{
		"* ",
		"**",
	}
	got := alignRight(input, 4)
	want := []string{
		"  * ",
		"  **",
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("alignRight line %d = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestAlignCenter(t *testing.T) {
	input := []string{
		"* ",
		"**",
	}
	got := alignCenter(input, 5)
	want := []string{
		" *   ",
		" **  ",
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("alignCenter line %d = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestAlignJustify(t *testing.T) {
	charMap, err := banner.LoadBannerFile("banners/standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}

	input := "a b"
	lines := converter.ConvertText(charMap, input)
	got := alignJustify(lines, input, charMap, 30)

	if len(got) != len(lines) {
		t.Fatalf("expected %d lines, got %d", len(lines), len(got))
	}

	for _, line := range got {
		if len(line) != 30 {
			t.Fatalf("expected justified width 30, got %q", line)
		}
	}
}

func TestAlignJustifyCompressesExtraSpacesInInput(t *testing.T) {
	charMap, err := banner.LoadBannerFile("banners/standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}

	width := 180
	inputA := "hello pre ty fce"
	inputB := "hello   pre   ty   fce"

	gotA := alignJustify(converter.ConvertText(charMap, inputA), inputA, charMap, width)
	gotB := alignJustify(converter.ConvertText(charMap, inputB), inputB, charMap, width)

	if !reflect.DeepEqual(gotA, gotB) {
		t.Fatalf("justify output should ignore extra spaces between words.\nA: %#v\nB: %#v", gotA, gotB)
	}

	for _, line := range gotA {
		if len(line) != width {
			t.Fatalf("expected justified width %d, got %q", width, line)
		}
	}
}

func TestAlignJustifyReturnsOriginalWhenNoSpace(t *testing.T) {
	charMap, err := banner.LoadBannerFile("banners/standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}

	input := "hello world"
	lines := converter.ConvertText(charMap, input)
	width := 5

	got := alignJustify(lines, input, charMap, width)
	if len(got) != len(lines) {
		t.Fatalf("expected %d lines when width is too small, got %d", len(lines), len(got))
	}

	expectedWidth := 0
	for _, word := range strings.Fields(input) {
		for _, char := range word {
			glyph := charMap[char]
			if len(glyph) > 0 {
				expectedWidth += len(glyph[0])
			}
		}
	}

	for i := range lines {
		if len(got[i]) != expectedWidth {
			t.Fatalf("expected line %d width %d, got %d", i, expectedWidth, len(got[i]))
		}
	}
}

func TestAlignLongLineKeepsOriginal(t *testing.T) {
	input := []string{
		"1234567890",
	}
	got := alignLeft(input, 5)
	if got[0] != input[0] {
		t.Fatalf("expected original line unchanged, got %q want %q", got[0], input[0])
	}
}
