package lab

import (
	"fmt"
	"os"
)

// FileReporter prints test details to a file
type FileReporter struct {
	file *os.File
}

// OutputReporter prints test details to STDOUT
var OutputReporter *FileReporter = &FileReporter{os.Stdout}

// BuildFileReporter instantiate a FileReporter for the provided file
func BuildFileReporter(file *os.File) *FileReporter {
	return &FileReporter{file}
}

// Asserted does nothing
func (reporter *FileReporter) Asserted() {}

// ContextEntered prints the context name
func (reporter *FileReporter) ContextEntered(prose string) {
	fmt.Fprintf(reporter.file, "%s\n", prose)
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
	fmt.Fprintf(reporter.file, "\t%s\n", prose)
}
