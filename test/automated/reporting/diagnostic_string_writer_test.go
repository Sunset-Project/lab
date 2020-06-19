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
			text1 := "A line\n"
			text2 := "Another line\n"
			fmt.Fprintf(writer, text1)
			fmt.Fprintf(writer, text2)

			writes := writer.LastRecorded(2)

			Test("Records all written strings", func() {
				Assert(writes[0] == text1)
				Assert(writes[1] == text2)
			})
		})
	})
}
