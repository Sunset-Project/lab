package reporting

import (
	"fmt"
	"io"
	"os"

	"github.com/sunset-project/lab/asserting"
	"github.com/sunset-project/lab/trace"
)

// IOReporter prints test details to a device
type IOReporter struct {
	Device io.Writer
}

var stdoutReporter IOReporter = IOReporter{Device: os.Stdout}

// StdoutReporter prints test details to STDOUT
func StdoutReporter() IOReporter { return stdoutReporter }

// Asserted does nothing
func (reporter IOReporter) Asserted() {}

// ContextEntered prints the context name
func (reporter IOReporter) ContextEntered(prose string) {
	fmt.Fprintf(reporter.Device, "%s\n", prose)
}

// ContextExited does nothing
func (reporter IOReporter) ContextExited(prose string) {}

// ContextSkipped does nothing
func (reporter IOReporter) ContextSkipped(prose string) {}

// ContextSucceeded does nothing
func (reporter IOReporter) ContextSucceeded(prose string) {}

// ContextFailed does nothing
func (reporter IOReporter) ContextFailed(prose string) {}

// PanicInvoked does nothing
func (reporter IOReporter) PanicInvoked(msg trace.Message) {
	if err, ok := msg.Data().(asserting.AssertionError); ok {
		fmt.Fprintf(reporter.Device, "\t%s\n", err.Error())
	} else {
		fmt.Fprintf(reporter.Device, "\t%s\n\t%+v\n", msg.Error(), msg.StackTrace())
	}
}

// TestFailed does nothing
func (reporter IOReporter) TestFailed(prose string) {}

// TestFinished does nothing
func (reporter IOReporter) TestFinished(prose string) {}

// TestPassed does nothing
func (reporter IOReporter) TestPassed(prose string) {}

// TestSkipped does nothing
func (reporter IOReporter) TestSkipped(prose string) {}

// TestStarted does nothing
func (reporter IOReporter) TestStarted(prose string) {
	fmt.Fprintf(reporter.Device, "\t%s\n", prose)
}
