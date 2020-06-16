package trace

import (
	"fmt"

	"github.com/pkg/errors"
)

// Message holds data returned from panic, with Stacktrace attached
type Message interface {
	Data() interface{}
	StackTrace() errors.StackTrace
	error
}

type message struct {
	data interface{}
	err  errorWithStackTrace
}

// Data provides argument passed to the panic function
func (msg message) Data() interface{} { return msg.data }

// DataString provides `Data` converted to `string` if possible, using `fmt.Stringer` or `Data` directly if already a `string`
func (msg message) DataString() (string, bool) {
	if stringer, ok := msg.data.(fmt.Stringer); ok {
		return stringer.String(), true
	}

	if text, ok := msg.data.(string); ok {
		return text, true
	}

	return "", false
}

// StackTrace provides the stack trace for the message
func (msg message) StackTrace() errors.StackTrace { return msg.err.StackTrace() }

// Error provides the error message for a panic
func (msg message) Error() string {
	text := "Panic"

	if dataText, ok := msg.DataString(); ok {
		text = fmt.Sprintf("Panic: %s", dataText)
	}

	return text
}

// NewMessage instantiates a new `Message` with `errors.StackTrace` attached
func NewMessage(data interface{}) Message {
	err := errors.New("Panic")
	errWithStack := err.(errorWithStackTrace)
	msg := message{data: data, err: errWithStack}

	return Message(msg)
}

type errorWithStackTrace interface {
	StackTrace() errors.StackTrace
	error
}
