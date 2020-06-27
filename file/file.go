// Package file provides utility to read file contents
package file

import (
	"os"
)

// Lines represents a single file line with a buffer indicating the lines before and after it
type Lines struct {
	FirstLine int
	Before    []string
	Exact     string
	After     []string
}

// ReadLineWithBuffers returns the specified line with a buffer of lines before and after it
func ReadLineWithBuffers(path string, currentLine int, beforeLinesAmount int, afterLinesAmount int) (Lines, error) {
	lastLine := currentLine + afterLinesAmount
	firstLine := currentLine - beforeLinesAmount
	if firstLine < 0 {
		firstLine = 1
	}

	lines := Lines{FirstLine: firstLine}
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return lines, err
	}

	scanner := NewLinesScanner(file)
	lineIndex := 1

	for scanner.Next() {
		line, err := scanner.Get()

		if err != nil {
			return lines, err
		}

		if lineIndex >= firstLine && lineIndex < currentLine {
			lines.Before = append(lines.Before, line)
		}
		if lineIndex == currentLine {
			lines.Exact = line
		}
		if lineIndex > currentLine && lineIndex <= lastLine {
			lines.After = append(lines.After, line)
		}

		lineIndex++
	}

	return lines, nil
}
