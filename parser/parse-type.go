package parser

import (
	"fmt"
	"spl/compiler"
	"spl/lexer"
	"spl/logs"
)

func (p *Parser) parseType() compiler.VariableType {
	t := p.expect(lexer.IDENTIFIER, fmt.Sprintf("Syntax Error: expected type but got %s", logs.TokenToString(p.at())))

	return compiler.VariableType{Type: t.Symbol}
}
