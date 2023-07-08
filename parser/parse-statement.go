package parser

import (
	"spl/compiler"
	"spl/lexer"
)

func (p *Parser) parseStatement() compiler.Node {
	switch p.at().Type {
	case lexer.LET:
		return p.parseVariable()
	case lexer.EOF:
		return nil
	default:
		expression := p.parseExpression()
		p.expectNewLine()

		return expression
	}
}
