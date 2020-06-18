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
	Recorded              []DiagnosticMessage
}

// DiagnosticMessage represents a call a `Reporter` method along with the provided arguments
type DiagnosticMessage struct {
	Signal Signal
	Data   interface{}
}

// Signal is used to identify what method has been called on the `DiagnosticReporter`
type Signal string

// Signals registered by the `DiagnosticReporter`
const (
	SigAsserted         Signal = "SigAsserted"
	SigContextEntered   Signal = "SigContextEntered"
	SigContextExited    Signal = "SigContextExited"
	SigContextSkipped   Signal = "SigContextSkipped"
	SigContextSucceeded Signal = "SigContextSucceeded"
	SigContextFailed    Signal = "SigContextFailed"
	SigPanicInvoked     Signal = "SigPanicInvoked"
	SigTestFailed       Signal = "SigTestFailed"
	SigTestFinished     Signal = "SigTestFinished"
	SigTestPassed       Signal = "SigTestPassed"
	SigTestSkipped      Signal = "SigTestSkipped"
	SigTestStarted      Signal = "SigTestStarted"
)

// LastRecorded returns the last N recorded diagnostic messages
func (reporter *DiagnosticReporter) LastRecorded(amount uint) []DiagnosticMessage {
	if amount <= 0 {
		return []DiagnosticMessage{}
	}

	size := len(reporter.Recorded)

	return reporter.Recorded[size-int(amount):]
}

// LastRecordedSignals returns the last N recorded signals
func (reporter *DiagnosticReporter) LastRecordedSignals(amount uint) []Signal {
	signals := []Signal{}
	recorded := reporter.LastRecorded(amount)

	for _, message := range recorded {
		signals = append(signals, message.Signal)
	}

	return signals
}

// Asserted increases AssertedCount by 1
func (reporter *DiagnosticReporter) Asserted() {
	reporter.AssertedCount = reporter.AssertedCount + 1
	message := DiagnosticMessage{SigAsserted, nil}
	reporter.Recorded = append(reporter.Recorded, message)
}

// ContextEntered increases ContextEnteredCount by 1
func (reporter *DiagnosticReporter) ContextEntered(prose string) {
	reporter.ContextEnteredCount = reporter.ContextEnteredCount + 1
	message := DiagnosticMessage{SigContextEntered, prose}
	reporter.Recorded = append(reporter.Recorded, message)
}

// ContextExited increases ContextExitedCount by 1
func (reporter *DiagnosticReporter) ContextExited(prose string) {
	reporter.ContextExitedCount = reporter.ContextExitedCount + 1
	message := DiagnosticMessage{SigContextExited, prose}
	reporter.Recorded = append(reporter.Recorded, message)
}

// ContextSkipped increases ContextSkippedCount by 1
func (reporter *DiagnosticReporter) ContextSkipped(prose string) {
	reporter.ContextSkippedCount = reporter.ContextSkippedCount + 1
	message := DiagnosticMessage{SigContextSkipped, prose}
	reporter.Recorded = append(reporter.Recorded, message)
}

// ContextSucceeded increases ContextSucceededCount by 1
func (reporter *DiagnosticReporter) ContextSucceeded(prose string) {
	reporter.ContextSucceededCount = reporter.ContextSucceededCount + 1
	message := DiagnosticMessage{SigContextSucceeded, prose}
	reporter.Recorded = append(reporter.Recorded, message)
}

// ContextFailed increases ContextFailedCount by 1
func (reporter *DiagnosticReporter) ContextFailed(prose string) {
	reporter.ContextFailedCount = reporter.ContextFailedCount + 1
	message := DiagnosticMessage{SigContextFailed, prose}
	reporter.Recorded = append(reporter.Recorded, message)
}

// PanicInvoked increases PanicInvokedCount by 1
func (reporter *DiagnosticReporter) PanicInvoked(traceMsg trace.Message) {
	reporter.PanicInvokedCount = reporter.PanicInvokedCount + 1
	message := DiagnosticMessage{SigPanicInvoked, traceMsg}
	reporter.Recorded = append(reporter.Recorded, message)
}

// TestFailed increases TestFailedCount by 1
func (reporter *DiagnosticReporter) TestFailed(prose string) {
	reporter.TestFailedCount = reporter.TestFailedCount + 1
	message := DiagnosticMessage{SigTestFailed, prose}
	reporter.Recorded = append(reporter.Recorded, message)
}

// TestFinished increases TestFinishedCount by 1
func (reporter *DiagnosticReporter) TestFinished(prose string) {
	reporter.TestFinishedCount = reporter.TestFinishedCount + 1
	message := DiagnosticMessage{SigTestFinished, prose}
	reporter.Recorded = append(reporter.Recorded, message)
}

// TestPassed increases TestPassedCount by 1
func (reporter *DiagnosticReporter) TestPassed(prose string) {
	reporter.TestPassedCount = reporter.TestPassedCount + 1
	message := DiagnosticMessage{SigTestPassed, prose}
	reporter.Recorded = append(reporter.Recorded, message)
}

// TestSkipped increases TestSkippedCount by 1
func (reporter *DiagnosticReporter) TestSkipped(prose string) {
	reporter.TestSkippedCount = reporter.TestSkippedCount + 1
	message := DiagnosticMessage{SigTestSkipped, prose}
	reporter.Recorded = append(reporter.Recorded, message)
}

// TestStarted increases TestStartedCount by 1
func (reporter *DiagnosticReporter) TestStarted(prose string) {
	reporter.TestStartedCount = reporter.TestStartedCount + 1
	message := DiagnosticMessage{SigTestStarted, prose}
	reporter.Recorded = append(reporter.Recorded, message)
}
