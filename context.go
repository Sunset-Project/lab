// Package lab is a testing framework
package lab

import (
	"fmt"
	"runtime"
	// "testing"
)

var TestFiles = make(map[string]interface{})

// Context do stuff
func Context(args ...interface{}) {
	switch len(args) {
	case 3:
		msg := args[0].(string)
		// t := args[1].(*testing.T)
		// context := args[2].(func())
		_, file, _, ok := runtime.Caller(1)

		if !ok {
			panic("Failed determining compiled file")
		}

		TestFiles[file] = msg
		fmt.Printf("Storing %s = %s\n", file, msg)
	// case 2:
	default:
		panic("Invalid number of arguments for Context")
	}
	fmt.Printf("%#+v\n", TestFiles)
}
