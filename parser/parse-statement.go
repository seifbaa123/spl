package parser

import (
	"spl/lexer"
	"spl/node"
)

func (p *Parser) parseStatement() node.Node {
	switch p.At().Type {
	case lexer.EOF:
		return nil
	default:
		expression := p.parseExpression()
		p.ExpectNewLine()

		return expression
	}
}
