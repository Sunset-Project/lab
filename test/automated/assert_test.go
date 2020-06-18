package lab_test

import (
	"testing"

	"github.com/sunset-project/lab"
	"github.com/sunset-project/lab/asserting"
	"github.com/sunset-project/lab/controls"
	"github.com/sunset-project/lab/reporting"
)

func TestAssert(t *testing.T) {
	Context, Test, Assert := lab.StartSession(t)
	Context("Assert", func() {
		session, reporter, controller := controls.SessionExample()

		Context("True", func() {
			assert := session.Assertion()

			assert(true)
			signals := reporter.LastRecordedSignals(1)

			Test("Report sequence is Asserted", func() {
				Assert(signals[0] == reporting.SigAsserted)
			})
		})

		Context("False", func() {
			assert := session.Assertion()
			immediateFailures := controller.ImmediateFailures

			session.Test("", func() { panic(nil) })

			Test("Panics with AssertionError", func() {
				Assert.PanicMsg(
					func() { assert(false) },
					func(err interface{}) bool {
						_, ok := err.(asserting.AssertionError)
						return ok
					},
				)
			})

			Test("Test controller Fails Now", func() {
				Assert(controller.ImmediateFailures == immediateFailures+1)
			})

			signals := reporter.LastRecordedSignals(1)

			Test("Report sequence is Asserted", func() {
				Assert(signals[0] == reporting.SigAsserted)
			})
		})
	})
}
