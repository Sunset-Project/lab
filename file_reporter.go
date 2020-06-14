package lab

import (
	"fmt"
	"os"
)

// FileReporter prints test details to a file
type FileReporter struct {
	File *os.File
}

// StdoutReporter prints test details to STDOUT
var StdoutReporter *FileReporter = &FileReporter{File: os.Stdout}

// NewFileReporter instantiate a FileReporter for the provided file
func NewFileReporter(file *os.File) *FileReporter {
	return &FileReporter{File: file}
}

// Asserted does nothing
func (reporter *FileReporter) Asserted() {}

// ContextEntered prints the context name
func (reporter *FileReporter) ContextEntered(prose string) {
	fmt.Fprintf(reporter.File, "%s\n", prose)
}

// ContextExited does nothing
func (reporter *FileReporter) ContextExited(prose string) {}

// ContextSkipped does nothing
func (reporter *FileReporter) ContextSkipped(prose string) {}

// ContextSucceeded does nothing
func (reporter *FileReporter) ContextSucceeded(prose string) {}

// ContextFailed does nothing
func (reporter *FileReporter) ContextFailed(prose string) {}

// PanicInvoked does nothing
func (reporter *FileReporter) PanicInvoked(msg interface{}) {}

// TestFailed does nothing
func (reporter *FileReporter) TestFailed(prose string) {}

// TestFinished does nothing
func (reporter *FileReporter) TestFinished(prose string) {}

// TestPassed does nothing
func (reporter *FileReporter) TestPassed(prose string) {}

// TestSkipped does nothing
func (reporter *FileReporter) TestSkipped(prose string) {}

// TestStarted does nothing
func (reporter *FileReporter) TestStarted(prose string) {
	fmt.Fprintf(reporter.File, "\t%s\n", prose)
}
