package main

import (
	"bufio"
	"fmt"
	"os"
	"proglan/lexicalanalizer"
)

func main() {
	var reader *bufio.Reader
	readFromFile := true
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Print("> ")
		reader = bufio.NewReader(os.Stdin)
		readFromFile = false
	} else {
		reader = bufio.NewReader(input)
		defer input.Close()
	}

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		tokens, err := lexicalanalizer.GetTokens(line)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Tokens:")
			for _, token := range tokens {
				fmt.Printf("Type: %s Value: %s\n", token.TokenType, token.Value)
			}
		}
		if !readFromFile {
			fmt.Print("> ")
		}
	}
}
