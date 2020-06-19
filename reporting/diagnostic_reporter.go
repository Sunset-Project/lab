package reporting

import (
	"github.com/sunset-project/lab/trace"
)

// DiagnosticReporter provides affordances to inspect Reporter usage
type DiagnosticReporter struct {
	Recorded []DiagnosticMessage
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
	SigAsserted         Signal = "Asserted"
	SigContextEntered   Signal = "ContextEntered"
	SigContextExited    Signal = "ContextExited"
	SigContextSkipped   Signal = "ContextSkipped"
	SigContextSucceeded Signal = "ContextSucceeded"
	SigContextFailed    Signal = "ContextFailed"
	SigPanicInvoked     Signal = "PanicInvoked"
	SigTestFailed       Signal = "TestFailed"
	SigTestFinished     Signal = "TestFinished"
	SigTestPassed       Signal = "TestPassed"
	SigTestSkipped      Signal = "TestSkipped"
	SigTestStarted      Signal = "TestStarted"
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

// Asserted records SigAsserted
func (reporter *DiagnosticReporter) Asserted() {
	message := DiagnosticMessage{SigAsserted, nil}
	reporter.Recorded = append(reporter.Recorded, message)
}

// ContextEntered records SigContextEntered
func (reporter *DiagnosticReporter) ContextEntered(prose string) {
	message := DiagnosticMessage{SigContextEntered, prose}
	reporter.Recorded = append(reporter.Recorded, message)
}

// ContextExited records SigContextExited
func (reporter *DiagnosticReporter) ContextExited(prose string) {
	message := DiagnosticMessage{SigContextExited, prose}
	reporter.Recorded = append(reporter.Recorded, message)
}

// ContextSkipped records SigContextSkipped
func (reporter *DiagnosticReporter) ContextSkipped(prose string) {
	message := DiagnosticMessage{SigContextSkipped, prose}
	reporter.Recorded = append(reporter.Recorded, message)
}

// ContextSucceeded records SigContextSucceeded
func (reporter *DiagnosticReporter) ContextSucceeded(prose string) {
	message := DiagnosticMessage{SigContextSucceeded, prose}
	reporter.Recorded = append(reporter.Recorded, message)
}

// ContextFailed records SigContextFailed
func (reporter *DiagnosticReporter) ContextFailed(prose string) {
	message := DiagnosticMessage{SigContextFailed, prose}
	reporter.Recorded = append(reporter.Recorded, message)
}

// PanicInvoked records SigPanicInvoked
func (reporter *DiagnosticReporter) PanicInvoked(traceMsg trace.Message) {
	message := DiagnosticMessage{SigPanicInvoked, traceMsg}
	reporter.Recorded = append(reporter.Recorded, message)
}

// TestFailed records SigTestFailed
func (reporter *DiagnosticReporter) TestFailed(prose string) {
	message := DiagnosticMessage{SigTestFailed, prose}
	reporter.Recorded = append(reporter.Recorded, message)
}

// TestFinished records SigTestFinished
func (reporter *DiagnosticReporter) TestFinished(prose string) {
	message := DiagnosticMessage{SigTestFinished, prose}
	reporter.Recorded = append(reporter.Recorded, message)
}

// TestPassed records SigTestPassed
func (reporter *DiagnosticReporter) TestPassed(prose string) {
	message := DiagnosticMessage{SigTestPassed, prose}
	reporter.Recorded = append(reporter.Recorded, message)
}

// TestSkipped records SigTestSkipped
func (reporter *DiagnosticReporter) TestSkipped(prose string) {
	message := DiagnosticMessage{SigTestSkipped, prose}
	reporter.Recorded = append(reporter.Recorded, message)
}

// TestStarted records SigTestStarted
func (reporter *DiagnosticReporter) TestStarted(prose string) {
	message := DiagnosticMessage{SigTestStarted, prose}
	reporter.Recorded = append(reporter.Recorded, message)
}
