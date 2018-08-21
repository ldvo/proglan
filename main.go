package main

import (
	"bufio"
	"fmt"
	"os"
	"proglan/lexicalanalizer"
	"proglan/syntacticalanalizer"
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
			if err := syntacticalanalizer.ParseExpression(tokens); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Valid expression.")
			}
		}
		if !readFromFile {
			fmt.Print("> ")
		}
	}
}
