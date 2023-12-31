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

	case lexer.PLUS_PLUS, lexer.MINUS_MINUS:
		operator := p.eat()
		identifier := p.expect(
			lexer.IDENTIFIER,
			fmt.Sprintf("SyntaxError: Expected identifier after %s but got %s", operator.Symbol, logs.TokenToString(p.at())),
		)

		expression = &compiler.Increment{
			Operator:   operator,
			Identifier: identifier,
		}

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

	if p.at().Type == lexer.QUESTION_MARK {
		expression = p.parseTernaryOperator(expression)
	}

	return expression
}
