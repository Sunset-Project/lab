package reporting_test

import (
	"fmt"
	"testing"

	"github.com/sunset-project/lab"
	"github.com/sunset-project/lab/reporting"
)

func TestDiagnosticStringWriter(t *testing.T) {
	Context, Test, Assert := lab.StartSession(t)
	Context("Diagnostic String Writer", func() {
		writer := &reporting.DiagnosticStringWriter{}

		Context("Writes", func() {
			text := "A line\n"
			otherText := "Another line\n"
			fmt.Fprintf(writer, text)
			fmt.Fprintf(writer, otherText)

			recorded := writer.LastRecorded(2)

			Test("Records all written strings", func() {
				Assert(recorded[0] == text)
				Assert(recorded[1] == otherText)
			})
		})
	})
}
