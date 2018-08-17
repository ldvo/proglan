package main

import (
	"bufio"
	"fmt"
	"os"
	"proglan/lexicalanalizer"
)

func main() {
	input, err := os.Open("input.txt")
	defer input.Close()
	if err != nil {
		fmt.Print(err)
		return
	}

	reader := bufio.NewReader(input)
	data, _, err := reader.ReadLine()
	if err != nil {
		fmt.Print(err)
		return
	}
	tokens, err := lexicalanalizer.GetTokens(data)
	if err != nil {
		fmt.Print(err)
		return
	}
	for _, token := range tokens {
		fmt.Printf("Type: %s \nValue: %s\n", token.TokenType, token.Value)
	}
}
