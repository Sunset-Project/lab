// Package sgr provides utilities to apply text styling in the terminal
package sgr

import (
	"fmt"
)

// Code is a numerical representation of a certain text style in the terminal
type Code uint

// SGR codes
const (
	Reset          Code = 0
	Bold           Code = 1
	Faint          Code = 2
	Italic         Code = 3
	Underline      Code = 4
	ResetIntensity Code = 22
	ResetItalic    Code = 23
	ResetUnderline Code = 24
	Black          Code = 30
	Red            Code = 31
	Green          Code = 32
	Yellow         Code = 33
	Blue           Code = 34
	Magenta        Code = 35
	Cyan           Code = 36
	White          Code = 37
	ResetFg        Code = 39
	BlackBg        Code = 40
	RedBg          Code = 41
	GreenBg        Code = 42
	YellowBg       Code = 43
	BlueBg         Code = 44
	MagentaBg      Code = 45
	CyanBg         Code = 46
	WhiteBg        Code = 47
	ResetBg        Code = 49
)

// String converts code into a string
func (code Code) String() string {
	return fmt.Sprintf("%d", code)
}
