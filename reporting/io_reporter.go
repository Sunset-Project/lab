package reporting

import (
	"fmt"
	"io"
	"os"
)

// IOReporter prints test details to a device
type IOReporter struct {
	Device io.Writer
}

// StdoutReporter prints test details to STDOUT
var StdoutReporter *IOReporter = &IOReporter{Device: os.Stdout}

// NewIOReporter instantiate a FileReporter for the provided device
func NewIOReporter(device io.Writer) *IOReporter {
	return &IOReporter{Device: device}
}

// Asserted does nothing
func (reporter *IOReporter) Asserted() {}

// ContextEntered prints the context name
func (reporter *IOReporter) ContextEntered(prose string) {
	fmt.Fprintf(reporter.Device, "%s\n", prose)
}

// ContextExited does nothing
func (reporter *IOReporter) ContextExited(prose string) {}

// ContextSkipped does nothing
func (reporter *IOReporter) ContextSkipped(prose string) {}

// ContextSucceeded does nothing
func (reporter *IOReporter) ContextSucceeded(prose string) {}

// ContextFailed does nothing
func (reporter *IOReporter) ContextFailed(prose string) {}

// PanicInvoked does nothing
func (reporter *IOReporter) PanicInvoked(msg interface{}) {}

// TestFailed does nothing
func (reporter *IOReporter) TestFailed(prose string) {}

// TestFinished does nothing
func (reporter *IOReporter) TestFinished(prose string) {}

// TestPassed does nothing
func (reporter *IOReporter) TestPassed(prose string) {}

// TestSkipped does nothing
func (reporter *IOReporter) TestSkipped(prose string) {}

// TestStarted does nothing
func (reporter *IOReporter) TestStarted(prose string) {
	fmt.Fprintf(reporter.Device, "\t%s\n", prose)
}
