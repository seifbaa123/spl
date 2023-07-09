package parser

import (
	"spl/compiler"
	"spl/lexer"
)

func (p *Parser) parseStatement() compiler.Node {
	switch p.at().Type {
	case lexer.LET:
		st := p.parseVariable()
		p.expectNewLine()
		return st

	case lexer.CONST:
		st := p.parseConstant()
		p.expectNewLine()
		return st

	case lexer.IDENTIFIER:
		if p.next().Type == lexer.EQUALS {
			st := p.parseReassignVariable()
			p.expectNewLine()
			return st
		} else {
			expression := p.parseExpression()
			p.expectNewLine()
			return expression
		}

	case lexer.EOF:
		return nil

	default:
		expression := p.parseExpression()
		p.expectNewLine()
		return expression
	}
}
