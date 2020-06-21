package reporting_test

import (
	"testing"

	"github.com/sunset-project/lab"
	"github.com/sunset-project/lab/reporting"
)

func TestIOReporter(t *testing.T) {
	Context, Test, Assert := lab.StartSession(t)
	Context("IO Reporter", func() {
		writer := &reporting.DiagnosticStringWriter{}
		reporter := &reporting.IOReporter{Device: writer}

		Context("Asserted", func() {
			reporter.Asserted()
			written := len(writer.Recorded)

			Test("Nothing is written", func() {
				Assert(written == 0)
			})
		})
	})
}
