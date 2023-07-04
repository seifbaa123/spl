package parser

import (
	"spl/lexer"
	"spl/statements"
)

func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{
		Tokens: tokens,
	}
}

func (p *Parser) ProduceAst() *statements.Program {
	var program statements.Program
	for p.At().Type != lexer.EOF {
		if p.At().Type == lexer.END_OF_LINE {
			p.Eat()
		}

		node := p.parseStatement()
		if node != nil {
			program.Body = append(program.Body, node)
		}
	}

	return &program
}
