package lab

// AssertionError is used as panic message when assertion fails
type AssertionError struct {
	Msg  string
	File string
	Line int
}

// Error provides the error message for a failing assertion
func (err *AssertionError) Error() string {
	if err.Msg == "" {
		return "Assertion failed"
	}

	return err.Msg
}
