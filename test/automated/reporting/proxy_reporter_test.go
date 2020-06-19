package lab_test

import (
	"testing"

	"github.com/sunset-project/lab"
	"github.com/sunset-project/lab/reporting"
	"github.com/sunset-project/lab/trace"
)

func TestProxyReporter(t *testing.T) {
	Context, Test, Assert := lab.StartSession(t)
	Context("Proxy Reporter", func() {
		diagnosticReporter1 := &reporting.DiagnosticReporter{}
		diagnosticReporter2 := &reporting.DiagnosticReporter{}
		reporter := reporting.ProxyReporter{diagnosticReporter1, diagnosticReporter2}

		Context("Broadcasts Asserted to all reporters", func() {
			reporter.Asserted()

			signals1 := diagnosticReporter1.LastRecordedSignals(1)
			signals2 := diagnosticReporter2.LastRecordedSignals(1)

			Test("Asserted is recorded by all diagnostic reporters", func() {
				Assert(signals1[0] == reporting.SigAsserted)
				Assert(signals2[0] == reporting.SigAsserted)
			})
		})

		Context("Broadcasts ContextEntered to all reporters", func() {
			reporter.ContextEntered("ContextEntered")

			signals1 := diagnosticReporter1.LastRecordedSignals(1)
			signals2 := diagnosticReporter2.LastRecordedSignals(1)

			Test("ContextEntered is recorded by all diagnostic reporters", func() {
				Assert(signals1[0] == reporting.SigContextEntered)
				Assert(signals2[0] == reporting.SigContextEntered)
			})
		})

		Context("Broadcasts ContextSkipped to all reporters", func() {
			reporter.ContextSkipped("ContextSkipped")

			signals1 := diagnosticReporter1.LastRecordedSignals(1)
			signals2 := diagnosticReporter2.LastRecordedSignals(1)

			Test("ContextSkipped is recorded by all diagnostic reporters", func() {
				Assert(signals1[0] == reporting.SigContextSkipped)
				Assert(signals2[0] == reporting.SigContextSkipped)
			})
		})

		Context("Broadcasts ContextSucceeded to all reporters", func() {
			reporter.ContextSucceeded("ContextSucceeded")

			signals1 := diagnosticReporter1.LastRecordedSignals(1)
			signals2 := diagnosticReporter2.LastRecordedSignals(1)

			Test("ContextSucceeded is recorded by all diagnostic reporters", func() {
				Assert(signals1[0] == reporting.SigContextSucceeded)
				Assert(signals2[0] == reporting.SigContextSucceeded)
			})
		})

		Context("Broadcasts ContextFailed to all reporters", func() {
			reporter.ContextFailed("ContextFailed")

			signals1 := diagnosticReporter1.LastRecordedSignals(1)
			signals2 := diagnosticReporter2.LastRecordedSignals(1)

			Test("ContextFailed is recorded by all diagnostic reporters", func() {
				Assert(signals1[0] == reporting.SigContextFailed)
				Assert(signals2[0] == reporting.SigContextFailed)
			})
		})

		Context("Broadcasts ContextExited to all reporters", func() {
			reporter.ContextExited("ContextExited")

			signals1 := diagnosticReporter1.LastRecordedSignals(1)
			signals2 := diagnosticReporter2.LastRecordedSignals(1)

			Test("ContextExited is recorded by all diagnostic reporters", func() {
				Assert(signals1[0] == reporting.SigContextExited)
				Assert(signals2[0] == reporting.SigContextExited)
			})
		})

		Context("Broadcasts PanicInvoked to all reporters", func() {
			traceMsg := trace.NewMessage(nil)
			reporter.PanicInvoked(traceMsg)

			signals1 := diagnosticReporter1.LastRecordedSignals(1)
			signals2 := diagnosticReporter2.LastRecordedSignals(1)

			Test("PanicInvoked is recorded by all diagnostic reporters", func() {
				Assert(signals1[0] == reporting.SigPanicInvoked)
				Assert(signals2[0] == reporting.SigPanicInvoked)
			})
		})

		Context("Broadcasts TestStarted to all reporters", func() {
			reporter.TestStarted("TestStarted")

			signals1 := diagnosticReporter1.LastRecordedSignals(1)
			signals2 := diagnosticReporter2.LastRecordedSignals(1)

			Test("TestStarted is recorded by all diagnostic reporters", func() {
				Assert(signals1[0] == reporting.SigTestStarted)
				Assert(signals2[0] == reporting.SigTestStarted)
			})
		})

		Context("Broadcasts TestSkipped to all reporters", func() {
			reporter.TestSkipped("TestSkipped")

			signals1 := diagnosticReporter1.LastRecordedSignals(1)
			signals2 := diagnosticReporter2.LastRecordedSignals(1)

			Test("TestSkipped is recorded by all diagnostic reporters", func() {
				Assert(signals1[0] == reporting.SigTestSkipped)
				Assert(signals2[0] == reporting.SigTestSkipped)
			})
		})

		Context("Broadcasts TestPassed to all reporters", func() {
			reporter.TestPassed("TestPassed")

			signals1 := diagnosticReporter1.LastRecordedSignals(1)
			signals2 := diagnosticReporter2.LastRecordedSignals(1)

			Test("TestPassed is recorded by all diagnostic reporters", func() {
				Assert(signals1[0] == reporting.SigTestPassed)
				Assert(signals2[0] == reporting.SigTestPassed)
			})
		})

		Context("Broadcasts TestFailed to all reporters", func() {
			reporter.TestFailed("TestFailed")

			signals1 := diagnosticReporter1.LastRecordedSignals(1)
			signals2 := diagnosticReporter2.LastRecordedSignals(1)

			Test("TestFailed is recorded by all diagnostic reporters", func() {
				Assert(signals1[0] == reporting.SigTestFailed)
				Assert(signals2[0] == reporting.SigTestFailed)
			})
		})

		Context("Broadcasts TestFinished to all reporters", func() {
			reporter.TestFinished("TestFinished")

			signals1 := diagnosticReporter1.LastRecordedSignals(1)
			signals2 := diagnosticReporter2.LastRecordedSignals(1)

			Test("TestFinished is recorded by all diagnostic reporters", func() {
				Assert(signals1[0] == reporting.SigTestFinished)
				Assert(signals2[0] == reporting.SigTestFinished)
			})
		})
	})
}
