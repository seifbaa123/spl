package main

import (
	"fmt"
	"os"
	"spl/lexer"
	"spl/utils"
)

func main() {
	src := utils.ReadSource()
	lexer := lexer.NewLexer(src, os.Args[1])
	fmt.Println(lexer.Tokenize())
}
