package reporting

import (
	"strings"
)

// OutputWriter generates a string to be written to `io.Writer` as output of test execution
type OutputWriter struct {
	StylingEnabled   bool
	IndentationDepth uint
	textBuffer       strings.Builder
	mode             outputWriterMode
}

type outputWriterMode uint

const (
	modeText outputWriterMode = iota
	modeEscapeSequence
)

// Writes text directly to the buffer
func (writer *OutputWriter) Write(text string) *OutputWriter {
	writer.textBuffer.WriteString(text)
	return writer
}

// IncreaseIndentation increases indentation depth by 1
func (writer *OutputWriter) IncreaseIndentation() *OutputWriter {
	writer.IndentationDepth++
	return writer
}

// DecreaseIndentation decreases indentation depth by 1
func (writer *OutputWriter) DecreaseIndentation() *OutputWriter {
	writer.IndentationDepth--
	return writer
}

// Compose returns the generated string and resets the buffer for the next one
func (writer *OutputWriter) Compose() string {
	output := writer.textBuffer.String()
	writer.textBuffer.Reset()

	return output
}

// Text appends the provided string to the buffer
func (writer *OutputWriter) Text(text string) *OutputWriter {
	if writer.mode == modeEscapeSequence {
		writer.mode = modeText
		writer.Write("m")
	}

	writer.Write(text)

	return writer
}

// Indent appends indentation space based on the current indentation depth
func (writer *OutputWriter) Indent() *OutputWriter {
	for depth := uint(0); depth <= writer.IndentationDepth; depth++ {
		writer.Text(" ")
	}
	return writer
}

// NewLine appends a new line character
func (writer *OutputWriter) NewLine() *OutputWriter {
	return writer.Text("\n")
}

// EscapeCode writes an escape sequence to style the text if styling is enabled
func (writer *OutputWriter) EscapeCode(code sgr.Code) *OutputWriter {
	if writer.mode == modeText && !writer.StylingEnabled {
		return writer
	}

	if writer.mode == modeText {
		writer.mode == modeEscapeSequence

		writer.Write("\e[")
	} else {
		writer.Write(";")
	}

	writer.Write(code.String())

	return writer
}
