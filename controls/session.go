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
