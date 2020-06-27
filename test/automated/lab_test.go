package lab_test

import (
	"testing"

	"github.com/sunset-project/lab"
	"github.com/sunset-project/lab/controls"
	"github.com/sunset-project/lab/reporting"
)

func TestLab(t *testing.T) {
	Context, Test, Assert := lab.StartSession(t)
	Context("Lab Session", func() {
		session, reporter, _ := controls.SessionExample()

		Context("Asserting in nested Context and Test", func() {
			Context("Assert False", func() {
				controls.AssertFalseInNestedContextTestExample(session)

				signals := reporter.LastRecordedSignals(8)

				Test("Report sequence is ContextEntered, TestStarted, Asserted, PanicInvoked, TestFailed, TestFinished, ContextSucceeded, ContextExited", func() {
					Assert(signals[0] == reporting.SigContextEntered)
					Assert(signals[1] == reporting.SigTestStarted)
					Assert(signals[2] == reporting.SigAsserted)
					Assert(signals[3] == reporting.SigPanicInvoked)
					Assert(signals[4] == reporting.SigTestFailed)
					Assert(signals[5] == reporting.SigTestFinished)
					Assert(signals[6] == reporting.SigContextSucceeded)
					Assert(signals[7] == reporting.SigContextExited)
				})
			})

			Context("Assert True", func() {
				controls.AssertTrueInNestedContextTestExample(session)

				signals := reporter.LastRecordedSignals(7)

				Test("Report sequence is ContextEntered, TestStarted, Asserted, TestPassed, TestFinished, ContextSucceeded, ContextExited", func() {
					Assert(signals[0] == reporting.SigContextEntered)
					Assert(signals[1] == reporting.SigTestStarted)
					Assert(signals[2] == reporting.SigAsserted)
					Assert(signals[3] == reporting.SigTestPassed)
					Assert(signals[4] == reporting.SigTestFinished)
					Assert(signals[5] == reporting.SigContextSucceeded)
					Assert(signals[6] == reporting.SigContextExited)
				})
			})
		})
	})
}
