package reporting

import (
	"fmt"
	"io"
	"os"

	"github.com/sunset-project/lab/asserting"
	"github.com/sunset-project/lab/sgr"
	"github.com/sunset-project/lab/trace"
)

// IOReporter prints test details to a device
type IOReporter struct {
	Device io.Writer
	output *OutputWriter
}

var stdoutReporter IOReporter = IOReporter{
	Device: os.Stdout,
	output: &OutputWriter{StylingEnabled: true},
}

// StdoutReporter prints test details to STDOUT
func StdoutReporter() IOReporter { return stdoutReporter }

// Asserted does nothing
func (reporter IOReporter) Asserted() {}

// ContextEntered prints the context name
func (reporter IOReporter) ContextEntered(prose string) {
	text := reporter.output.
		Indent().
		EscapeCode(sgr.Green).
		Text(prose).
		EscapeCode(sgr.ResetFg).
		NewLine().
		Flush()

	reporter.output.IncreaseIndentation()

	fmt.Fprintf(reporter.Device, text)
}

// ContextExited does nothing
func (reporter IOReporter) ContextExited(prose string, success bool) {
	if prose == "" {
		return
	}

	reporter.output.DecreaseIndentation()

	if reporter.output.IndentationDepth == 0 {
		reporter.output.NewLine()
	}

	text := reporter.output.Flush()

	fmt.Fprintf(reporter.Device, text)
}

// ContextSkipped does nothing
func (reporter IOReporter) ContextSkipped(prose string) {
	if prose == "" {
		return
	}
	reporter.output.Indent()

	if reporter.output.StylingEnabled {
		reporter.output.
			EscapeCode(sgr.Yellow).
			Text(prose).
			EscapeCode(sgr.ResetFg)
	} else {
		reporter.output.
			Text(prose).
			Text(" (skipped)")
	}

	reporter.output.NewLine()
	if reporter.output.IndentationDepth == 0 {
		reporter.output.NewLine()
	}

	text := reporter.output.Flush()
	fmt.Fprintf(reporter.Device, text)
}

// ContextSucceeded does nothing
func (reporter IOReporter) ContextSucceeded(prose string) {}

// ContextFailed does nothing
func (reporter IOReporter) ContextFailed(prose string) {}

// PanicInvoked does nothing
func (reporter IOReporter) PanicInvoked(msg trace.Message) {
	reporter.output.Indent()
	if err, ok := msg.Data().(asserting.AssertionError); ok {
		reporter.output.
			Text(err.Error()).
			NewLine()
	} else {
		stacktrace := fmt.Sprintf("%+v", msg.StackTrace())
		reporter.output.
			Text(msg.Error()).
			NewLine().
			Indent().
			Text(stacktrace).
			NewLine()
	}

	text := reporter.output.Flush()
	fmt.Fprintf(reporter.Device, text)
}

// TestFailed does nothing
func (reporter IOReporter) TestFailed(prose string) {}

// TestFinished does nothing
func (reporter IOReporter) TestFinished(prose string, success bool) {
	reporter.output.Indent()

	fgColor := sgr.Green
	if !success {
		reporter.output.EscapeCode(sgr.Bold)
		fgColor = sgr.Red
	}

	title := prose
	if title == "" {
		title = "Test"
	}

	reporter.output.
		EscapeCode(fgColor).
		Text(title).
		EscapeCode(sgr.ResetFg)

	if !success {
		reporter.output.EscapeCode(sgr.ResetIntensity)
	}

	reporter.output.NewLine()
	text := reporter.output.Flush()
	fmt.Fprintf(reporter.Device, text)
}

// TestPassed does nothing
func (reporter IOReporter) TestPassed(prose string) {}

// TestSkipped does nothing
func (reporter IOReporter) TestSkipped(prose string) {
	title := prose
	if title == "" {
		title = "Test"
	}

	reporter.output.Indent()

	if reporter.output.StylingEnabled {
		reporter.output.
			EscapeCode(sgr.Yellow).
			Text(title).
			EscapeCode(sgr.ResetFg)
	} else {
		reporter.output.
			Text(title).
			Text(" (skipped)")
	}

	reporter.output.NewLine()

	text := reporter.output.Flush()
	fmt.Fprintf(reporter.Device, text)
}

// TestStarted does nothing
func (reporter IOReporter) TestStarted(prose string) {}
