package lab_test

import (
	"testing"

	"github.com/sunset-project/lab"
	"github.com/sunset-project/lab/labctrls"
)

func TestContext(t *testing.T) {
	Context, Test, Assert := lab.StartSession(t)
	Context("Context", func() {
		session, reporter, _ := labctrls.SessionExample()

		Context("Empty block", func() {
			skipped := reporter.ContextSkippedCount
			session.Context("")

			Test("Skipped reported", func() {
				Assert(reporter.ContextSkippedCount == skipped+1)
			})
		})
	})
}
