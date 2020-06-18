package lab_test

import (
	"testing"

	"github.com/sunset-project/lab"
	"github.com/sunset-project/lab/controls"
	"github.com/sunset-project/lab/reporting"
)

func TestTest(t *testing.T) {
	Context, Test, Assert := lab.StartSession(t)
	Context("Test", func() {
		session, reporter, controller := controls.SessionExample()

		Context("No function", func() {
			session.Test("")

			signals := reporter.LastRecordedSignals(3)

			Test("Report sequence is Start, Skip, Finish", func() {
				Assert(signals[0] == reporting.SigTestStarted)
				Assert(signals[1] == reporting.SigTestSkipped)
				Assert(signals[2] == reporting.SigTestFinished)
			})
		})

		Context("Function with panic", func() {
			immediateFailures := controller.ImmediateFailures

			session.Test("", func() { panic(nil) })

			Test("Test controller Fails Now", func() {
				Assert(controller.ImmediateFailures == immediateFailures+1)
			})

			signals := reporter.LastRecordedSignals(4)

			Test("Report sequence is Start, Panic, Fail, Finish", func() {
				Assert(signals[0] == reporting.SigTestStarted)
				Assert(signals[1] == reporting.SigPanicInvoked)
				Assert(signals[2] == reporting.SigTestFailed)
				Assert(signals[3] == reporting.SigTestFinished)
			})
		})

		Context("Function without panic", func() {
			immediateFailures := controller.ImmediateFailures

			session.Test("", func() {})

			Test("Test controller didn't Fail", func() {
				Assert(controller.ImmediateFailures == immediateFailures)
			})

			signals := reporter.LastRecordedSignals(3)

			Test("Report sequence is Start, Pass, Finish", func() {
				Assert(signals[0] == reporting.SigTestStarted)
				Assert(signals[1] == reporting.SigTestPassed)
				Assert(signals[2] == reporting.SigTestFinished)
			})
		})
	})
}
