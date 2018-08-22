package syntacticalanalizer

import (
	"fmt"
	"proglan/enums/tokentype"
	"proglan/lexicalanalizer"
)

// ParseExpression checks if a sequence of tokens is valid
func ParseExpression(tokens []lexicalanalizer.Token) error {
	if len(tokens) < 2 {
		return fmt.Errorf("Empty expression")
	}
	nextToken, err := parseParenthesis(tokens, 0)
	if err != nil {
		return err
	}
	if nextToken != len(tokens)-1 {
		return fmt.Errorf("Unexpected end of expression")
	}
	if tokens[nextToken].TokenType != tokentype.EOF {
		return fmt.Errorf("Expected '.' at the end")
	}
	return nil
}

func parseParenthesis(tokens []lexicalanalizer.Token, nextToken int) (int, error) {
	if tokens[nextToken].TokenType != tokentype.Parenthesis && tokens[nextToken].Value != "(" {
		return 0, fmt.Errorf("Expected '(' at position %d, found '%s'", nextToken, tokens[nextToken].Value)
	}

	var err error
	nextToken, err = parseOperation(tokens, nextToken+1)
	if err != nil {
		return 0, err
	}

	if tokens[nextToken].TokenType != tokentype.Parenthesis && tokens[nextToken].Value != ")" {
		return 0, fmt.Errorf("Expected ')' at position %d, found '%s'", nextToken, tokens[nextToken].Value)
	}
	nextToken++
	return nextToken, nil
}

func parseOperation(tokens []lexicalanalizer.Token, nextToken int) (int, error) {
	var err error
	if tokens[nextToken].TokenType != tokentype.Operator {
		return 0, fmt.Errorf("Expected operator at position %d, found '%s'", nextToken, tokens[nextToken].Value)
	}
	nextToken++

	nextToken, err = parseOperationTerm(tokens, nextToken)
	if err != nil {
		return 0, err
	}

	nextToken, err = parseOperationTerm(tokens, nextToken)
	if err != nil {
		return 0, err
	}
	return nextToken, nil
}

func parseOperationTerm(tokens []lexicalanalizer.Token, nextToken int) (int, error) {
	if tokens[nextToken].TokenType == tokentype.Number || tokens[nextToken].TokenType == tokentype.Variable {
		nextToken++
	} else {
		var err error
		nextToken, err = parseParenthesis(tokens, nextToken)
		if err != nil {
			return 0, err
		}
	}
	return nextToken, nil
}
