package panic

import (
	"fmt"

	"github.com/pkg/errors"
)

// Message holds data returned from panic, with Stacktrace attached
type Message struct {
	Data interface{}
}

// Error provides the error message for a panic
func (msg Message) Error() string {
	text := "Panic"

	if stringer, ok := msg.Data.(fmt.Stringer); ok {
		text = fmt.Sprintf("Panic: %s", stringer.String())
	}

	return text
}

// NewMessage instantiates a new `Message` with StackTrace attached
func NewMessage() Message {
	msg := Message{}
	errors.WithStack(msg)

	return msg
}
