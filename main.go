package main

import (
	"ascii-art/internal/banner"
	"ascii-art/internal/converter"
	"fmt"
	"os"
)

// main parses CLI input, loads the banner, and prints plain or colored ASCII art.
func main() {
	opts, err := parseArgs(os.Args)
	if err != nil {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println("EX: go run . --color=<color> <substring to be colored> \"something\"")
		return
	}

	// Validate input contains only printable ASCII characters (32-126)
	for _, char := range opts.Input {
		if char < 32 || char > 126 {
			fmt.Println("Please only use : printable ASCII characters 32 to 126 (use man ascii if needed)")
			return
		}
	}

	// Load banner file containing ASCII art templates
	bannerPath := "banners/" + opts.Banner + ".txt"
	charMap, err := banner.LoadBannerFile(bannerPath)
	if err != nil {
		fmt.Println("Error loading", bannerPath)
		return
	}

	if opts.UseColor {
		colorCode, ok := colorCodeFromName(opts.Color)
		if !ok {
			fmt.Println("Unsupported color:", opts.Color)
			fmt.Println("Supported colors: black, red, green, yellow, blue, magenta, cyan, white")
			return
		}
		art := converter.ConvertTextWithColor(charMap, opts.Input, opts.Substring, colorCode)
		for _, line := range art {
			fmt.Println(line)
		}
		return
	}

	// Convert input text to ASCII art
	art := converter.ConvertText(charMap, opts.Input)
	if opts.OutputFile != "" {
		err := writeOutputToFile(opts, art)
		if err != nil {
			fmt.Println("Error writing output to file:", err)
		} else {
			fmt.Println("Output written to", opts.OutputFile)
		}
		return
	}
	if opts.Align != "" {
		width := getTerminalWidth()
		switch opts.Align {
		case "left":
			art = alignLeft(art, width)
		case "right":
			art = alignRight(art, width)
		case "center":
			art = alignCenter(art, width)
		case "justify":
			art = alignJustify(art, opts.Input, charMap, width)
		default:
			fmt.Println("Unsupported alignment:", opts.Align)
			fmt.Println("Supported alignments: left, right, center, justify")
			return
		}
	}
	for _, line := range art {
		fmt.Println(line)
	}

}
