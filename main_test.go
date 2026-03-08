package main

import (
	"ascii-art/internal/banner"
	"ascii-art/internal/converter"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestValidString(t *testing.T) {
	input := "Hello"
	expected := []string{
		" _    _          _   _          ",
		"| |  | |        | | | |         ",
		"| |__| |   ___  | | | |   ___   ",
		"|  __  |  / _ \\ | | | |  / _ \\  ",
		"| |  | | |  __/ | | | | | (_) | ",
		"|_|  |_|  \\___| |_| |_|  \\___/  ",
		"                                ",
		"                                ",
	}

	charMap, err := banner.LoadBannerFile("banners/standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}

	result := converter.ConvertText(charMap, input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected:\n%v\n\nGot:\n%v", expected, result)
	}
}

func TestArgumentParsing(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"single string", []string{"cmd", "Hello"}, false},
		{"color whole string", []string{"cmd", "--color=red", "Hello"}, false},
		{"color substring", []string{"cmd", "--color=red", "kit", "a king kitten have kit"}, false},
		{"output file", []string{"cmd", "--output=filename.txt", "Hello", "shadow"}, false},
		{"empty output filename", []string{"cmd", "--output=", "Hello", "shadow"}, true},
		{"bad output flag format", []string{"cmd", "--output", "Hello", "shadow"}, true},
		{"output filename without txt suffix", []string{"cmd", "--output=file", "Hello", "shadow"}, true},
		{"align left", []string{"cmd", "--align=left", "Hello", "standard"}, false},
		{"align right", []string{"cmd", "--align=right", "Hello"}, false},
		{"align center", []string{"cmd", "--align=center", "Hello"}, false},
		{"align justify", []string{"cmd", "--align=justify", "Hello"}, false},
		{"invalid align mode", []string{"cmd", "--align=middle", "Hello"}, true},
		{"bad align flag format", []string{"cmd", "--align", "Hello"}, true},
		{"empty align mode", []string{"cmd", "--align=", "Hello", "shadow"}, true},
		{"too few args", []string{"cmd"}, true},
		{"bad flag name", []string{"cmd", "--colour=red", "Hello"}, true},
		{"bad flag format", []string{"cmd", "--color", "red", "Hello"}, true},
		{"empty color", []string{"cmd", "--color=", "Hello"}, true},
		{"too many args", []string{"cmd", "--color=red", "a", "b", "c"}, true},
		{"string with standard banner", []string{"cmd", "hello", "standard"}, false},
		{"string with shadow banner", []string{"cmd", "hello", "shadow"}, false},
		{"string with thinkertoy banner", []string{"cmd", "hello", "thinkertoy"}, false},
		{"color whole string with standard banner", []string{"cmd", "--color=red", "hello", "standard"}, false},
		{"color whole string with shadow banner", []string{"cmd", "--color=red", "hello", "shadow"}, false},
		{"color whole string with thinkertoy banner", []string{"cmd", "--color=red", "hello", "thinkertoy"}, false},
		{"color substring with standard banner", []string{"cmd", "--color=red", "kit", "hello world", "standard"}, false},
		{"color substring with thinkertoy banner", []string{"cmd", "--color=red", "kit", "hello world", "thinkertoy"}, false},
		{"unknown banner", []string{"cmd", "hello", "unknown"}, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			opts, err := parseArgs(tc.args)
			if (err != nil) != tc.wantErr {
				t.Fatalf("wantErr=%v, got err=%v", tc.wantErr, err)
			}
			if !tc.wantErr {
				// Ensure all expected fields are set correctly
				if opts.Input == "" {
					t.Error("Expected non-empty Input")
				}
				if opts.UseColor && opts.Color == "" {
					t.Error("Expected non-empty Color when UseColor is true")
				}
			}
		})
	}
}

func TestUnsupportedColor(t *testing.T) {
	colorCode, supported := colorCodeFromName("orange")
	if supported {
		t.Errorf("Expected unsupported color to return false, got true with code %s", colorCode)
	}
	if colorCode != "" {
		t.Errorf("Expected unsupported color to return empty code, got %s", colorCode)
	}
}

func TestDifferentBannersProduceDifferentOutput(t *testing.T) {
	input := "A"

	standardMap, err := banner.LoadBannerFile("banners/standard.txt")
	if err != nil {
		t.Fatalf("Failed to load standard banner: %v", err)
	}

	shadowMap, err := banner.LoadBannerFile("banners/shadow.txt")
	if err != nil {
		t.Fatalf("Failed to load shadow banner: %v", err)
	}

	standardResult := converter.ConvertText(standardMap, input)
	shadowResult := converter.ConvertText(shadowMap, input)

	if reflect.DeepEqual(standardResult, shadowResult) {
		t.Fatal("expected different output for different banners")
	}
}

func TestOutputFileWrite(t *testing.T) {
	tempDir := t.TempDir()
	outputFile := filepath.Join(tempDir, "test_output.txt")
	input := "Test"

	charMap, err := banner.LoadBannerFile("banners/standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}

	expectedLines := converter.ConvertText(charMap, input)
	expectedContent := ""
	for i, line := range expectedLines {
		expectedContent += line
		if i < len(expectedLines)-1 {
			expectedContent += "\n"
		}
	}

	opts := cliOptions{
		Input:      input,
		OutputFile: outputFile,
		Banner:     "standard",
	}

	err = writeOutputToFile(opts, expectedLines)
	if err != nil {
		t.Fatalf("Failed to write output to file: %v", err)
	}

	data, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	if string(data) != expectedContent {
		t.Fatalf("Expected file content:\n%q\nGot:\n%q", expectedContent, string(data))
	}
}
