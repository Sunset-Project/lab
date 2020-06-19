package lab_test

import (
	"testing"

	"github.com/sunset-project/lab"
	"github.com/sunset-project/lab/reporting"
	"github.com/sunset-project/lab/trace"
)

func TestDiagnosticReporter(t *testing.T) {
	Context, Test, Assert := lab.StartSession(t)
	Context("Diagnostic Reporter", func() {
		reporter := &reporting.DiagnosticReporter{}

		Context("Messages", func() {
			reporter.ContextEntered("entered")
			reporter.ContextExited("exited")

			messages := reporter.LastRecorded(2)

			Test("Recorded in ascending order", func() {
				Assert(messages[0].Signal == reporting.SigContextEntered)
				Assert(messages[0].Data.(string) == "entered")
				Assert(messages[1].Signal == reporting.SigContextExited)
				Assert(messages[1].Data.(string) == "exited")
			})

			signals := reporter.LastRecordedSignals(2)

			Test("Signals correspond to messages' signals", func() {
				Assert(messages[0].Signal == signals[0])
				Assert(messages[1].Signal == signals[1])
			})
		})

		Context("Asserted", func() {
			reporter.Asserted()

			messages := reporter.LastRecorded(1)

			Test("Records Asserted signal", func() {
				Assert(messages[0].Signal == reporting.SigAsserted)
			})
		})

		Context("Context Entered", func() {
			reporter.ContextEntered("entered")

			messages := reporter.LastRecorded(1)

			Test("Records ContextEntered message", func() {
				Assert(messages[0].Signal == reporting.SigContextEntered)
				Assert(messages[0].Data.(string) == "entered")
			})
		})

		Context("Context Skipped", func() {
			reporter.ContextSkipped("skipped")

			messages := reporter.LastRecorded(1)

			Test("Records ContextSkipped message", func() {
				Assert(messages[0].Signal == reporting.SigContextSkipped)
				Assert(messages[0].Data.(string) == "skipped")
			})
		})

		Context("Context Succeeded", func() {
			reporter.ContextSucceeded("succeeded")

			messages := reporter.LastRecorded(1)

			Test("Records ContextSucceeded message", func() {
				Assert(messages[0].Signal == reporting.SigContextSucceeded)
				Assert(messages[0].Data.(string) == "succeeded")
			})
		})

		Context("Context Failed", func() {
			reporter.ContextFailed("failed")

			messages := reporter.LastRecorded(1)

			Test("Records ContextFailed message", func() {
				Assert(messages[0].Signal == reporting.SigContextFailed)
				Assert(messages[0].Data.(string) == "failed")
			})
		})

		Context("Context Exited", func() {
			reporter.ContextExited("exited")

			messages := reporter.LastRecorded(1)

			Test("Records ContextExited message", func() {
				Assert(messages[0].Signal == reporting.SigContextExited)
				Assert(messages[0].Data.(string) == "exited")
			})
		})

		Context("Panic Invoked", func() {
			traceMsg := trace.NewMessage(nil)
			reporter.PanicInvoked(traceMsg)

			messages := reporter.LastRecorded(1)

			Test("Records PanicInvoked message", func() {
				Assert(messages[0].Signal == reporting.SigPanicInvoked)
				Assert(messages[0].Data == traceMsg)
			})
		})

		Context("Test Started", func() {
			reporter.TestStarted("started")

			messages := reporter.LastRecorded(1)

			Test("Records TestStarted message", func() {
				Assert(messages[0].Signal == reporting.SigTestStarted)
				Assert(messages[0].Data.(string) == "started")
			})
		})

		Context("Test Skipped", func() {
			reporter.TestSkipped("skipped")

			messages := reporter.LastRecorded(1)

			Test("Records TestSkipped message", func() {
				Assert(messages[0].Signal == reporting.SigTestSkipped)
				Assert(messages[0].Data.(string) == "skipped")
			})
		})

		Context("Test Passed", func() {
			reporter.TestPassed("passed")

			messages := reporter.LastRecorded(1)

			Test("Records TestPassed message", func() {
				Assert(messages[0].Signal == reporting.SigTestPassed)
				Assert(messages[0].Data.(string) == "passed")
			})
		})

		Context("Test Failed", func() {
			reporter.TestFailed("failed")

			messages := reporter.LastRecorded(1)

			Test("Records TestFailed message", func() {
				Assert(messages[0].Signal == reporting.SigTestFailed)
				Assert(messages[0].Data.(string) == "failed")
			})
		})

		Context("Test Finished", func() {
			reporter.TestFinished("finished")

			messages := reporter.LastRecorded(1)

			Test("Records TestFinished message", func() {
				Assert(messages[0].Signal == reporting.SigTestFinished)
				Assert(messages[0].Data.(string) == "finished")
			})
		})
	})
}
