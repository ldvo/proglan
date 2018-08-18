//go:generate stringer -type=RuneType

package tokentype

type TokenType int

const (
	Number TokenType = iota
	Variable
	Parenthesis
	Operator
	EOF
)
