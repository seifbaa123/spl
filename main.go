package main

import (
	"fmt"
	"os"
	"spl/lexer"
	"spl/parser"
	"spl/utils"
)

func main() {
	src := utils.ReadSource()
	lexer := lexer.NewLexer(src, os.Args[1])
	parser := parser.NewParser(lexer.Tokenize())
	fmt.Println(parser.ProduceAst())
}
