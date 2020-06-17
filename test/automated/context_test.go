package lab_test

import (
	"testing"

	"github.com/sunset-project/lab"
	"github.com/sunset-project/lab/controls"
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
