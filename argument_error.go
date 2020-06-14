package lab

import "fmt"

// ArgumentError is reported when passed argument is invalid or when arity is incorrect
type ArgumentError struct {
	Name string
	Msg  string
}

// Error provides the error message for invalid argument
func (err ArgumentError) Error() string {
	msg := "is invalid"
	if err.Msg != "" {
		msg = err.Msg
	}

	return fmt.Sprintf("`%s` %s", err.Name, msg)
}
