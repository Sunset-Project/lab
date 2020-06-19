package reporting

import "strings"

// DiagnosticStringWriter holds a record of each `Write` operation as a slice that can be inspected for diagnostic purposes
type DiagnosticStringWriter struct {
	Recorded []string
}

// Write records a new entry in the recorded strings
func (writer *DiagnosticStringWriter) Write(p []byte) (int, error) {
	var buf strings.Builder
	n, err := buf.Write(p)

	writer.Recorded = append(writer.Recorded, buf.String())

	return n, err
}

// LastRecorded returns the last N recorded strings
func (writer *DiagnosticStringWriter) LastRecorded(amount uint) []string {
	if amount <= 0 {
		return []string{}
	}

	size := len(writer.Recorded)

	return writer.Recorded[size-int(amount):]
}
