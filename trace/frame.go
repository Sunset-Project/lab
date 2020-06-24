package trace

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Entry represents one frame in the stacktrace from `Message`
type Entry struct{ Frame errors.Frame }

// FunctionName extracts source file path from an `errors.Frame`
func (entry Entry) FunctionName() string {
	frame := entry.Frame
	text := fmt.Sprintf("%+s", frame)
	funcNameAndPath := strings.SplitN(text, "\n\t", 2)
	return funcNameAndPath[0]
}

// SourcePath extracts source file path from an `errors.Frame`
func (entry Entry) SourcePath() string {
	frame := entry.Frame
	text := fmt.Sprintf("%+s", frame)
	funcNameAndPath := strings.SplitN(text, "\n\t", 2)
	return funcNameAndPath[1]
}

// SourceLine extracts source line from an `errors.Frame`
func (entry Entry) SourceLine() int {
	frame := entry.Frame
	text := fmt.Sprintf("%d", frame)
	line, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return line
}

// ShortFunctionName extracts source file path from an `errors.Frame`
func (entry Entry) ShortFunctionName() string {
	frame := entry.Frame
	return fmt.Sprintf("%n", frame)
}
