package parser

import (
	"fmt"
	"spl/compiler"
	"spl/lexer"
	"spl/logs"
)

func (p *Parser) parseTernaryOperator(expression compiler.Node) *compiler.TernaryOperator {
	token := p.expect(
		lexer.QUESTION_MARK,
		fmt.Sprintf("SyntaxError: expected ? but got %s", logs.TokenToString(p.at())),
	)

	ifExpression := p.parseExpression()

	p.expect(
		lexer.COLON,
		fmt.Sprintf("SyntaxError: expected : but got %s", logs.TokenToString(p.at())),
	)

	elseExpression := p.parseExpression()

	return &compiler.TernaryOperator{
		Token:          token,
		Condition:      expression,
		IfExpression:   ifExpression,
		ElseExpression: elseExpression,
	}
}
