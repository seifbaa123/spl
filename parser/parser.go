package parser

import (
	"spl/compiler"
	"spl/lexer"
)

func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{
		Tokens: tokens,
	}
}

func (p *Parser) ProduceAst() *compiler.Program {
	var program compiler.Program
	for p.at().Type != lexer.EOF {
		if p.at().Type == lexer.END_OF_LINE {
			p.eat()
		}

		node := p.parseStatement()
		if node != nil {
			program.Body = append(program.Body, node)
		}
	}

	return &program
}
