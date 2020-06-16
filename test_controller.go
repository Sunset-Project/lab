package lab

// TestController provides an API to halt a test in case of failure
type TestController interface {
	FailNow()
}
