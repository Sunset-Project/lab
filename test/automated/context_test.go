package lab_test

import (
	"testing"

	"github.com/sunset-project/lab"
	"github.com/sunset-project/lab/controls"
	"github.com/sunset-project/lab/reporting"
)

func TestContext(t *testing.T) {
	Context, Test, Assert := lab.StartSession(t)
	Context("Context", func() {
		session, reporter, controller := controls.SessionExample()

		Context("No function", func() {
			skipped := reporter.ContextSkippedCount

			session.Context("")

			Test("Skipped reported", func() {
				Assert(reporter.ContextSkippedCount == skipped+1)
			})

			signals := reporter.LastRecordedSignals(3)

			Test("Report sequence is Enter, Skip, Exit", func() {
				Assert(signals[0] == reporting.SigContextEntered)
				Assert(signals[1] == reporting.SigContextSkipped)
				Assert(signals[2] == reporting.SigContextExited)
			})
		})

		Context("Function with panic", func() {
			failed := reporter.ContextFailedCount
			immediateFailures := controller.ImmediateFailures

			session.Context("", func() { panic(nil) })

			Test("Failed reported", func() {
				Assert(reporter.ContextFailedCount == failed+1)
			})
			Test("Test controller fails", func() {
				Assert(controller.ImmediateFailures == immediateFailures+1)
			})
		})
	})
}
