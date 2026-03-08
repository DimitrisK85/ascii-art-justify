package main

import (
	"fmt"
	"os"
)

func writeOutputToFile(opts cliOptions, art []string) error {
	content := ""
	for i, line := range art {
		content += line
		if i < len(art)-1 {
			content += "\n"
		}
	}
	err := os.WriteFile(opts.OutputFile, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write output to file: %w", err)

	}
	return nil
}
