package main

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func alignLeft(lines []string, width int) []string {
	blockWidth := len(lines[0]) //getMaxLineWidth(lines)
	if blockWidth >= width {
		return append([]string{}, lines...)
	}

	aligned := make([]string, 0, len(lines))
	for _, line := range lines {
		aligned = append(aligned, line+strings.Repeat(" ", width-len(line)))
	}
	return aligned
}

func alignRight(lines []string, width int) []string {
	blockWidth := len(lines[0]) //getMaxLineWidth(lines)
	if blockWidth >= width {
		return append([]string{}, lines...)
	}

	leftPadding := width - blockWidth
	aligned := make([]string, 0, len(lines))
	for _, line := range lines {
		aligned = append(aligned, strings.Repeat(" ", leftPadding)+line+strings.Repeat(" ", blockWidth-len(line)))
	}
	return aligned
}

func alignCenter(lines []string, width int) []string {
	blockWidth := len(lines[0]) //getMaxLineWidth(lines)
	if blockWidth >= width {
		return append([]string{}, lines...)
	}

	leftPadding := (width - blockWidth) / 2
	aligned := make([]string, 0, len(lines))
	for _, line := range lines {
		content := strings.Repeat(" ", leftPadding) + line + strings.Repeat(" ", blockWidth-len(line))
		aligned = append(aligned, content+strings.Repeat(" ", width-len(content)))
	}
	return aligned
}

func alignJustify(lines []string, input string, charMap map[rune][]string, width int) []string {
	// Split input text into words to identify gaps for justify spacing.
	words := strings.Fields(input)
	if len(words) <= 1 {
		return alignLeft(lines, width)
	}

	// Keep behavior simple and safe when width is invalid or no space handling exists.
	//if width <= 0 {
	//	return append([]string{}, lines...)
	//}

	// Calculate the base width of each rendered word.
	totalWordWidth := 0
	for _, word := range words {
		for _, char := range word {
			glyph := charMap[char]
			if len(glyph) == 0 {
				continue
			}
			totalWordWidth += len(glyph[0]) //glyphWidth(glyph)
		}
	}

	gapCount := len(words) - 1
	totalSpaces := width - totalWordWidth
	if width < totalWordWidth {
		width = totalWordWidth
		totalSpaces = 0
	}
	//if gapCount <= 0 || totalSpaces < 0 {
	//	return append([]string{}, lines...)
	//}

	// Total spaces to distribute between words.
	spacePerGap := totalSpaces / gapCount
	extraSpaces := totalSpaces % gapCount

	// ASCII art is 8 lines tall for supported banners.
	result := make([]string, 8)
	for row := 0; row < 8; row++ {
		var line strings.Builder
		for i, word := range words {
			for _, char := range word {
				glyph := charMap[char]
				if len(glyph) <= row {
					continue
				}
				line.WriteString(glyph[row])
			}
			if i < len(words)-1 {
				line.WriteString(strings.Repeat(" ", spacePerGap))
				if i < extraSpaces {
					line.WriteString(" ")
				}
			}
		}
		result[row] = line.String()
	}

	return result
}

//func glyphWidth(glyph []string) int {
//	width := 0
//	for _, row := range glyph {
//		if len(row) > width {
//			width = len(row)
//		}
//	}
//	return width
//}

func getTerminalWidth() int {
	columns := os.Getenv("COLUMNS")
	if columns != "" {
		if width, err := strconv.Atoi(columns); err == nil && width > 0 {
			return width
		}
	}

	// Execute stty size command to get terminal dimensions (rows columns)
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err == nil {
		parts := strings.Fields(string(out))
		if len(parts) == 2 {
			if width, err := strconv.Atoi(parts[1]); err == nil && width > 0 {
				return width
			}
		}
	}

	// Fallback width for non-interactive environments
	return 80
}

//func getMaxLineWidth(lines []string) int {
//	maxWidth := 0
//	for _, line := range lines {
//		if len(line) > maxWidth {
//			maxWidth = len(line)
//		}
//	}
//	return maxWidth
//}
