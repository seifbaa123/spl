package parser

import (
	"fmt"
	"spl/compiler"
	"spl/lexer"
)

func (p *Parser) parseType() compiler.VariableType {
	t := p.expect(lexer.IDENTIFIER, fmt.Sprintf("Syntax Error: expected type but got %s", p.at().Symbol))

	return compiler.VariableType{Type: t.Symbol}
}
