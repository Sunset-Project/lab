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
			session.Context("")

			signals := reporter.LastRecordedSignals(3)

			Test("Report sequence is Enter, Skip, Exit", func() {
				Assert(signals[0] == reporting.SigContextEntered)
				Assert(signals[1] == reporting.SigContextSkipped)
				Assert(signals[2] == reporting.SigContextExited)
			})
		})

		Context("Function with panic", func() {
			immediateFailures := controller.ImmediateFailures

			session.Context("", func() { panic(nil) })

			Test("Test controller Fails Now", func() {
				Assert(controller.ImmediateFailures == immediateFailures+1)
			})

			signals := reporter.LastRecordedSignals(4)

			Test("Report sequence is Enter, Panic, Fail, Exit", func() {
				Assert(signals[0] == reporting.SigContextEntered)
				Assert(signals[1] == reporting.SigPanicInvoked)
				Assert(signals[2] == reporting.SigContextFailed)
				Assert(signals[3] == reporting.SigContextExited)
			})
		})
	})
}
