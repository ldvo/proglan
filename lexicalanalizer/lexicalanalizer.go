package lexicalanalizer

import (
	"fmt"
	"proglan/enums/runetype"
	"proglan/enums/tokentype"
)

const errorState = 999

// Token type
type Token struct {
	TokenType tokentype.TokenType
	Value     string
}

var transitionMatrix = []map[runetype.RuneType]int{
	// State 0
	{
		runetype.Invalid:     errorState,
		runetype.Digit:       1,
		runetype.Letter:      errorState,
		runetype.Pound:       2,
		runetype.Parenthesis: int(tokentype.Parenthesis) + 100,
		runetype.Operator:    int(tokentype.Operator) + 100,
		runetype.EOF:         int(tokentype.EOF) + 100,
	},
	// State 1
	{
		runetype.Invalid:     errorState,
		runetype.Digit:       1,
		runetype.Letter:      errorState,
		runetype.Pound:       int(tokentype.Number) + 100,
		runetype.Parenthesis: int(tokentype.Number) + 100,
		runetype.Operator:    int(tokentype.Number) + 100,
		runetype.EOF:         int(tokentype.Number) + 100,
	},
	// State 2
	{
		runetype.Invalid:     errorState,
		runetype.Digit:       errorState,
		runetype.Letter:      3,
		runetype.Pound:       errorState,
		runetype.Parenthesis: errorState,
		runetype.Operator:    errorState,
		runetype.EOF:         errorState,
	},
	// State 3
	{
		runetype.Invalid:     errorState,
		runetype.Digit:       3,
		runetype.Letter:      3,
		runetype.Pound:       int(tokentype.Variable) + 100,
		runetype.Parenthesis: int(tokentype.Variable) + 100,
		runetype.Operator:    int(tokentype.Variable) + 100,
		runetype.EOF:         int(tokentype.Variable) + 100,
	},
}

// GetTokens : Gets the tokens for a given byte array
func GetTokens(data string) ([]Token, error) {
	status := 0
	value := []rune{}
	tokens := []Token{}
	runes := []rune(data)
	for pos := 0; pos < len(runes); pos++ {
		runeType := runetype.GetRuneType(runes[pos])
		status = transitionMatrix[status][runeType]
		if status < 100 {
			// Keep reading
			value = append(value, runes[pos])
		} else if status < 200 {
			// Add token
			tokenType := tokentype.TokenType(status - 100)
			if tokenType == tokentype.Number || tokenType == tokentype.Variable {
				pos--
			} else {
				value = append(value, runes[pos])
			}
			token := Token{
				TokenType: tokenType,
				Value:     string(value),
			}
			tokens = append(tokens, token)
			status = 0
			value = []rune{}
		} else {
			// Invalid rune
			return tokens, fmt.Errorf("Syntax error at position %d, char '%s'", pos, string(runes[pos]))
		}
	}
	// If we were reading a number or a variable, add it
	if status == 1 {
		token := Token{
			TokenType: tokentype.Number,
			Value:     string(value),
		}
		tokens = append(tokens, token)
	}
	if status == 2 || status == 3 {
		token := Token{
			TokenType: tokentype.Variable,
			Value:     string(value),
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}
