package main

import (
	"os"
	"spl/compiler"
	"spl/lexer"
	"spl/parser"
	"spl/utils"
)

func main() {
	src := utils.ReadSource()
	lexer := lexer.NewLexer(src, os.Args[1])
	parser := parser.NewParser(lexer.Tokenize())
	ast := parser.ProduceAst()
	compiler.Compile(ast)
}
