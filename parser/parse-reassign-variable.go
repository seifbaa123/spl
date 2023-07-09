package parser

import (
	"fmt"
	"spl/compiler"
	"spl/lexer"
	"spl/logs"
)

func (p *Parser) parseReassignVariable() *compiler.ReassignVariable {
	variable := p.expect(lexer.IDENTIFIER, fmt.Sprintf("Syntax Error: expected variable name but got %s", logs.TokenToString(p.at())))
	op := p.expect(lexer.EQUALS, fmt.Sprintf("Syntax Error: expected = after variable name but got %s", logs.TokenToString(p.at())))

	expression := p.parseExpression()

	return &compiler.ReassignVariable{
		Op:         op,
		Variable:   variable,
		Expression: expression,
	}
}
