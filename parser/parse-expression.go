package parser

import (
	"spl/lexer"
	"spl/node"
)

func (p *Parser) parseExpression() node.Node {
	switch p.At().Type {
	case lexer.PRINT:
		return p.parsePrint()
	default:
		return p.parseAdding()
	}
}
