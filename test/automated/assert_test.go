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
		session, reporter, _ := controls.SessionExample()
		assert := session.Assertion()

		Context("True", func() {
			assert(true)
			signals := reporter.LastRecordedSignals(1)

			Test("Report sequence is Asserted", func() {
				Assert(signals[0] == reporting.SigAsserted)
			})
		})

		Context("False", func() {
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

			signals := reporter.LastRecordedSignals(1)

			Test("Report sequence is Asserted", func() {
				Assert(signals[0] == reporting.SigAsserted)
			})
		})

		Context("Panic", func() {
			Context("Function with no panic", func() {
				Test("Panics with AssertionError", func() {
					Assert.PanicMsg(
						func() { assert.Panic(func() {}) },
						func(err interface{}) bool {
							_, ok := err.(asserting.AssertionError)
							return ok
						},
					)
				})

				signals := reporter.LastRecordedSignals(1)

				Test("Report sequence is Asserted", func() {
					Assert(signals[0] == reporting.SigAsserted)
				})
			})

			Context("Function with panic", func() {
				assert.Panic(func() { panic(nil) })
				signals := reporter.LastRecordedSignals(1)

				Test("Report sequence is Asserted", func() {
					Assert(signals[0] == reporting.SigAsserted)
				})
			})
		})

		Context("Panic message", func() {
			Context("Function with no panic", func() {
				Test("Panics with AssertionError", func() {
					Assert.PanicMsg(
						func() { assert.Panic(func() {}) },
						func(err interface{}) bool {
							_, ok := err.(asserting.AssertionError)
							return ok
						},
					)
				})

				signals := reporter.LastRecordedSignals(1)

				Test("Report sequence is Asserted", func() {
					Assert(signals[0] == reporting.SigAsserted)
				})
			})
			Context("Function with panic", func() {
				Context("Message matches", func() {
					assert.PanicMsg(
						func() { panic(nil) },
						func(_ interface{}) bool { return true },
					)
					signals := reporter.LastRecordedSignals(1)

					Test("Report sequence is Asserted", func() {
						Assert(signals[0] == reporting.SigAsserted)
					})
				})

				Context("Message does not match", func() {
					panicNow := func() { panic(nil) }
					noMatch := func(_ interface{}) bool { return false }

					Test("Panics with AssertionError", func() {
						Assert.PanicMsg(
							func() { assert.PanicMsg(panicNow, noMatch) },
							func(err interface{}) bool {
								_, ok := err.(asserting.AssertionError)
								return ok
							},
						)
					})

					signals := reporter.LastRecordedSignals(1)

					Test("Report sequence is Asserted", func() {
						Assert(signals[0] == reporting.SigAsserted)
					})
				})
			})
		})
	})
}
