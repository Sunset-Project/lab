// Package trace provides utilities to trace data recovered from panics
package trace

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func DoStuff() {
	panic("whatever")
}

// FileLines returns the line in a file surrounded by N before lines and M after lines
func FileLines(path string, currentLine int, beforeLinesAmount int, afterLinesAmount int) ([]string, string, []string) {
	lastLine := currentLine + afterLinesAmount
	firstLine := currentLine - beforeLinesAmount
	if firstLine < 0 {
		firstLine = 1
	}

	var exactLine string
	beforeLines := []string{}
	afterLines := []string{}

	fmt.Println("readFileWithReadString")

	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	lineIndex := 1

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF || err == nil {
			if lineIndex >= firstLine && lineIndex < currentLine {
				beforeLines = append(beforeLines, line)
			}
			if lineIndex == currentLine {
				exactLine = line
			}
			if lineIndex > currentLine && lineIndex <= lastLine {
				afterLines = append(afterLines, line)
			}
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		lineIndex++
	}

	fmt.Printf("before: %+v\ncurrent: %+v\nafter: %+v\n", beforeLines, exactLine, afterLines)
	return beforeLines, exactLine, afterLines
}
