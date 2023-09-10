package parser

import (
	"spl/compiler"
	"spl/lexer"
)

func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{
		Tokens:          tokens,
		IsInParenthesis: false,
	}
}

func (p *Parser) ProduceAst() *compiler.Program {
	var program compiler.Program
	for p.at().Type != lexer.EOF {
		if p.at().Type == lexer.NEW_LINE {
			p.eat()
		}

		node := p.parseStatement()
		if node != nil {
			program.Body = append(program.Body, node)
		}
	}

	return &program
}
