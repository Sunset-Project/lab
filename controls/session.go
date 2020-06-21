package controls

import (
	"github.com/sunset-project/lab"
	"github.com/sunset-project/lab/reporting"
)

// SessionExample provides a `lab.Session` with diagnostic utilities
func SessionExample() (lab.Session, *reporting.DiagnosticReporter, *lab.DiagnosticTestController) {
	reporter := &reporting.DiagnosticReporter{}
	controller := &lab.DiagnosticTestController{}
	return lab.NewSession(controller, reporter), reporter, controller
}

// AssertFalseInNestedContextTestExample asserts value false
func AssertFalseInNestedContextTestExample(session lab.Session) {
	Context, Test, Assert := lab.UseSession(session)
	Context("Context", func() {
		Test("Test", func() {
			Assert(false)
		})
	})
}

// AssertTrueInNestedContextTestExample asserts value true
func AssertTrueInNestedContextTestExample(session lab.Session) {
	Context, Test, Assert := lab.UseSession(session)
	Context("Context", func() {
		Test("Test", func() {
			Assert(true)
		})
	})
}
