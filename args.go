package main

import (
	"fmt"
	"strings"
)

type cliOptions struct {
	Input      string
	Color      string
	Substring  string
	Banner     string
	OutputFile string
	Align      string
	UseColor   bool
}

func parseArgs(args []string) (cliOptions, error) {
	opts := cliOptions{Banner: "standard"}
	parseCount := len(args)

	if parseCount >= 3 && isValidBanner(args[parseCount-1]) {
		opts.Banner = args[parseCount-1]
		parseCount--
	}

	if parseCount == 2 {
		opts.Input = args[1]
		return opts, nil
	}

	if parseCount == 3 && strings.HasPrefix(args[1], "--color=") {
		color := strings.TrimPrefix(args[1], "--color=")
		if color == "" {
			return cliOptions{}, fmt.Errorf("empty color value")
		}
		opts.Input = args[2]
		opts.Color = color
		opts.UseColor = true
		return opts, nil
	}

	if parseCount == 3 && strings.HasPrefix(args[1], "--output=") {
		outputFile := strings.TrimPrefix(args[1], "--output=")
		if outputFile == "" {
			return cliOptions{}, fmt.Errorf("empty output filename")
		}
		if !strings.HasSuffix(outputFile, ".txt") {
			return cliOptions{}, fmt.Errorf("output filename must end with .txt")
		}
		opts.Input = args[2]
		opts.OutputFile = outputFile
		return opts, nil
	}

	if parseCount == 3 && strings.HasPrefix(args[1], "--align=") {
		align := strings.TrimPrefix(args[1], "--align=")
		if align == "" {
			return cliOptions{}, fmt.Errorf("empty align value")
		}
		if !isValidAlign(align) {
			return cliOptions{}, fmt.Errorf("invalid align value: %s", align)
		}
		opts.Input = args[2]
		opts.Align = align
		return opts, nil
	}

	if parseCount == 4 && strings.HasPrefix(args[1], "--color=") {
		color := strings.TrimPrefix(args[1], "--color=")
		if color == "" {
			return cliOptions{}, fmt.Errorf("empty color value")
		}
		opts.Input = args[3]
		opts.Color = color
		opts.Substring = args[2]
		opts.UseColor = true
		return opts, nil
	}

	return cliOptions{}, fmt.Errorf("invalid argument format")
}

func isValidBanner(name string) bool {
	switch name {
	case "standard", "shadow", "thinkertoy":
		return true
	default:
		return false
	}
}

func isValidAlign(align string) bool {
	switch align {
	case "left", "right", "center", "justify":
		return true
	default:
		return false
	}
}

// colorCodeFromName maps a supported color name to its ANSI escape code.
func colorCodeFromName(color string) (string, bool) {
	switch strings.ToLower(color) {
	case "black":
		return "\033[30m", true
	case "red":
		return "\033[31m", true
	case "green":
		return "\033[32m", true
	case "yellow":
		return "\033[33m", true
	case "blue":
		return "\033[34m", true
	case "magenta":
		return "\033[35m", true
	case "cyan":
		return "\033[36m", true
	case "white":
		return "\033[37m", true
	default:
		return "", false
	}
}
