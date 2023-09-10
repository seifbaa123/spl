package parser

import (
	"fmt"
	"spl/compiler"
	"spl/lexer"
	"spl/logs"
)

func (p *Parser) parseExpression() compiler.Node {
	var expression compiler.Node

	switch p.at().Type {
	case lexer.PRINT:
		expression = p.parsePrint()

	case lexer.OPEN_BRACKET:
		expression = p.parseList()

	default:
		for p.at().Type == lexer.NEW_LINE && p.IsInParenthesis {
			p.eat()
		}

		expression = p.parseLogical()
	}

	for p.at().Type == lexer.DOT {
		p.eat()
		expression = &compiler.PropertyExpression{
			Expression: expression,
			Property: p.expect(
				lexer.IDENTIFIER,
				fmt.Sprintf(
					"Syntax Error: expected property name but token %s",
					logs.TokenToString(p.at()),
				),
			),
		}
	}

	return expression
}
