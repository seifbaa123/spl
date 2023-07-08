package parser

import (
	"spl/compiler"
	"spl/lexer"
)

func (p *Parser) parseExpression() compiler.Node {
	switch p.at().Type {
	case lexer.PRINT:
		return p.parsePrint()
	default:
		return p.parseAdding()
	}
}
