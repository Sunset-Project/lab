package asserting

// Assertion provides a function to assert results, as well as a function to assert "comma ok" tuples and assert and recover from functions triggering panic
type Assertion func(...interface{})

// PanicMsg asserts that the provided function triggers panic with the provided message
func (assert Assertion) PanicMsg(do func(), assertMsg func(interface{}) bool) {
	defer func() {
		err := recover()

		assert(err != nil, "Panic expected")

		if err != nil {
			result := assertMsg(err)
			assert(result, "Invalid panic message")
		}
	}()
	do()
}

// Panic asserts that the provided function triggers panic
func (assert Assertion) Panic(do func()) {
	assert.PanicMsg(do, any)
}

func any(_ interface{}) bool { return true }
