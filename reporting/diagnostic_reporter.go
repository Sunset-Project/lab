package reporting

import (
	"github.com/sunset-project/lab/trace"
)

// DiagnosticReporter provides affordances to inspect Reporter usage
type DiagnosticReporter struct {
	AssertedCount         uint
	ContextEnteredCount   uint
	ContextExitedCount    uint
	ContextSkippedCount   uint
	ContextSucceededCount uint
	ContextFailedCount    uint
	PanicInvokedCount     uint
	TestFailedCount       uint
	TestFinishedCount     uint
	TestPassedCount       uint
	TestSkippedCount      uint
	TestStartedCount      uint
}

// Asserted increases AssertedCount by 1
func (reporter *DiagnosticReporter) Asserted() {
	reporter.AssertedCount = reporter.AssertedCount + 1
}

// ContextEntered increases ContextEnteredCount by 1
func (reporter *DiagnosticReporter) ContextEntered(prose string) {
	reporter.ContextEnteredCount = reporter.ContextEnteredCount + 1
}

// ContextExited increases ContextExitedCount by 1
func (reporter *DiagnosticReporter) ContextExited(prose string) {
	reporter.ContextExitedCount = reporter.ContextExitedCount + 1
}

// ContextSkipped increases ContextSkippedCount by 1
func (reporter *DiagnosticReporter) ContextSkipped(prose string) {
	reporter.ContextSkippedCount = reporter.ContextSkippedCount + 1
}

// ContextSucceeded increases ContextSucceededCount by 1
func (reporter *DiagnosticReporter) ContextSucceeded(prose string) {
	reporter.ContextSucceededCount = reporter.ContextSucceededCount + 1
}

// ContextFailed increases ContextFailedCount by 1
func (reporter *DiagnosticReporter) ContextFailed(prose string) {
	reporter.ContextFailedCount = reporter.ContextFailedCount + 1
}

// PanicInvoked increases PanicInvokedCount by 1
func (reporter *DiagnosticReporter) PanicInvoked(msg trace.Message) {
	reporter.PanicInvokedCount = reporter.PanicInvokedCount + 1
}

// TestFailed increases TestFailedCount by 1
func (reporter *DiagnosticReporter) TestFailed(prose string) {
	reporter.TestFailedCount = reporter.TestFailedCount + 1
}

// TestFinished increases TestFinishedCount by 1
func (reporter *DiagnosticReporter) TestFinished(prose string) {
	reporter.TestFinishedCount = reporter.TestFinishedCount + 1
}

// TestPassed increases TestPassedCount by 1
func (reporter *DiagnosticReporter) TestPassed(prose string) {
	reporter.TestPassedCount = reporter.TestPassedCount + 1
}

// TestSkipped increases TestSkippedCount by 1
func (reporter *DiagnosticReporter) TestSkipped(prose string) {
	reporter.TestSkippedCount = reporter.TestSkippedCount + 1
}

// TestStarted increases TestStartedCount by 1
func (reporter *DiagnosticReporter) TestStarted(prose string) {
	reporter.TestStartedCount = reporter.TestStartedCount + 1
}
