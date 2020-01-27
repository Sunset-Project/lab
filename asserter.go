package lab

// AssertionError is used as panic message when assertion fails
type AssertionError struct {
	Msg  string
	File string
	Line int
}

// Assertion provides a function to assert results, as well as a function to assert "comma ok" tuples and assert and recover from functions triggering panic
type Assertion func(...interface{})

// Error provides the error message for a failing assertion
func (err *AssertionError) Error() string {
	if err.Msg == "" {
		return "Assertion failed"
	}

	return err.Msg
}

// PanicMsg asserts that the provided function triggers panic with the provided message
func (assert Assertion) PanicMsg(assertMsg func(interface{}) bool, do func()) {
	defer func() {
		err := recover()

		assert(err != nil, "No panic")

		if err != nil {
			result := assertMsg(err)
			assert(result, "Invalid panic message")
		}
	}()
	do()
}

// Panic asserts that the provided function triggers panic
func (assert Assertion) Panic(do func()) {
	assert.PanicMsg(any, do)
}

// Ok asserts that the provided "comma ok" tuple is successful
func (assert Assertion) Ok(_ interface{}, ok bool) {
	assert(ok, "Ok tuple unsuccessful")
}

func any(_ interface{}) bool { return true }
