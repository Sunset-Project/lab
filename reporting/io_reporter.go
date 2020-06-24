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
	Device        io.Writer
	output        *OutputWriter
	previousError trace.Message
}

var stdoutReporter *IOReporter = &IOReporter{
	Device: os.Stdout,
	output: &OutputWriter{StylingEnabled: true},
}

// StdoutReporter prints test details to STDOUT
func StdoutReporter() *IOReporter { return stdoutReporter }

// Asserted does nothing
func (reporter *IOReporter) Asserted() {}

// ContextEntered prints the context name
func (reporter *IOReporter) ContextEntered(prose string) {
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
func (reporter *IOReporter) ContextExited(prose string, success bool) {
	if prose == "" {
		return
	}

	reporter.output.DecreaseIndentation()

	if reporter.output.IndentationDepth == 0 {
		reporter.output.NewLine()
	}

	if reporter.previousError != nil {
		reporter.output.IncreaseIndentation()
		reporter.PrintPreviousError()
		reporter.output.DecreaseIndentation()
	}

	text := reporter.output.Flush()

	fmt.Fprintf(reporter.Device, text)
}

// ContextSkipped does nothing
func (reporter *IOReporter) ContextSkipped(prose string) {
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
func (reporter *IOReporter) ContextSucceeded(prose string) {}

// ContextFailed does nothing
func (reporter *IOReporter) ContextFailed(prose string) {}

// PanicInvoked sets previous error
func (reporter *IOReporter) PanicInvoked(msg trace.Message) {
	reporter.previousError = msg
}

// PrintError outputs panic data with stacktrace
func (reporter *IOReporter) PrintError(msg trace.Message) {

	if err, ok := msg.Data().(asserting.AssertionError); ok {
		reporter.output.
			Indent().
			EscapeCode(sgr.Red).
			Text(err.Error()).
			EscapeCode(sgr.ResetFg).
			NewLine()
	} else {
		reporter.output.
			Indent().
			EscapeCode(sgr.White).
			EscapeCode(sgr.RedBg).
			Text(msg.Error()).
			EscapeCode(sgr.ResetBg).
			EscapeCode(sgr.ResetFg).
			NewLine().
			IncreaseIndentation().
			IncreaseIndentation()

		reporter.output.EscapeCode(sgr.Red)

		stackTrace := msg.StackTrace()
		isTopOmitted := false
		testRunnerIndex := len(stackTrace) - 3
		isFirstFrame := true

		for index, frame := range stackTrace {
			if index < 3 {
				continue
			}
			if index >= testRunnerIndex {
				break
			}

			if !isTopOmitted {
				isTopOmitted = true
				reporter.output.
					Indent().
					EscapeCode(sgr.Italic).
					Text("*omitted*").
					EscapeCode(sgr.ResetItalic).
					NewLine()
			}

			entry := trace.Entry{Frame: frame}

			funcName := entry.FunctionName()
			filePath := entry.SourcePath()
			line := entry.SourceLine()
			fileLine := fmt.Sprintf("%s:%d", filePath, line)

			if isFirstFrame {
				reporter.output.EscapeCode(sgr.Bold)
			}

			reporter.output.
				Indent().
				Text(funcName).
				NewLine().
				IncreaseIndentation().
				Indent().
				Text(fileLine).
				NewLine().
				DecreaseIndentation()

			if isFirstFrame {
				reporter.output.EscapeCode(sgr.ResetIntensity)
				isFirstFrame = false
			}
		}

		reporter.output.
			Indent().
			EscapeCode(sgr.Italic).
			Text("*omitted*").
			EscapeCode(sgr.ResetItalic).
			NewLine()

		reporter.output.
			DecreaseIndentation().
			DecreaseIndentation().
			EscapeCode(sgr.ResetFg).
			NewLine()
	}
}

// PrintPreviousError prints the last occurred error and resets it to nil
func (reporter *IOReporter) PrintPreviousError() {
	reporter.PrintError(reporter.previousError)
	reporter.previousError = nil
}

// TestFailed does nothing
func (reporter *IOReporter) TestFailed(prose string) {}

// TestFinished does nothing
func (reporter *IOReporter) TestFinished(prose string, success bool) {
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

	if reporter.previousError != nil {
		reporter.output.IncreaseIndentation()
		reporter.PrintPreviousError()
		reporter.output.DecreaseIndentation()
	}

	text := reporter.output.Flush()
	fmt.Fprintf(reporter.Device, text)
}

// TestPassed does nothing
func (reporter *IOReporter) TestPassed(prose string) {}

// TestSkipped does nothing
func (reporter *IOReporter) TestSkipped(prose string) {
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
func (reporter *IOReporter) TestStarted(prose string) {}
