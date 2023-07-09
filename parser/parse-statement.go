package parser

import (
	"spl/compiler"
	"spl/lexer"
)

func (p *Parser) parseStatement() compiler.Node {
	switch p.at().Type {
	case lexer.LET:
		let := p.parseVariable()
		p.expectNewLine()
		return let

	case lexer.CONST:
		constant := p.parseConstant()
		p.expectNewLine()
		return constant

	case lexer.EOF:
		return nil

	default:
		expression := p.parseExpression()
		p.expectNewLine()
		return expression
	}
}
