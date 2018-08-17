package runetype

import (
	"unicode"
)

// RuneType : The possible rune types
type RuneType int

const (
	// Invalid type
	Invalid RuneType = iota
	// Digit type
	Digit
	// Letter type
	Letter
	// Pound type
	Pound
	// Parenthesis type
	Parenthesis
	// Operator type
	Operator
	// EOF type
	EOF
)

// GetRuneType : Gets the rune type for a given rune
func GetRuneType(r rune) RuneType {
	if unicode.IsDigit(r) {
		return Digit
	}
	if unicode.IsLetter(r) {
		return Letter
	}
	switch r {
	case '#':
		return Pound
	case '(', ')':
		return Parenthesis
	case '+', '-', '*', '/':
		return Operator
	case '.':
		return EOF
	}
	return Invalid
}
