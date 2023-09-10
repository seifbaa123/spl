package parser

import (
	"fmt"
	"spl/compiler"
	"spl/lexer"
	"spl/logs"
)

func (p *Parser) parseType() *compiler.VariableType {
	t := p.expect(lexer.IDENTIFIER, fmt.Sprintf("Syntax Error: expected type but got %s", logs.TokenToString(p.at())))
	var subType *compiler.VariableType = nil

	if t.Symbol == "List" {
		p.expect(lexer.LESS, fmt.Sprintf("Syntax Error: expected < after List but got %s", logs.TokenToString(p.at())))
		subType = p.parseType()
		p.expect(lexer.GREATER, fmt.Sprintf("Syntax Error: expected > after subtype but got %s", logs.TokenToString(p.at())))
	}

	return &compiler.VariableType{Type: t.Symbol, SubType: subType}
}
