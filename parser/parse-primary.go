package parser

import (
	"fmt"
	"os"
	"spl/compiler"
	"spl/lexer"
	"spl/logs"
)

func (p *Parser) ParsePrimary() compiler.Node {
	var expression compiler.Node

	switch p.at().Type {
	case lexer.NUMBER:
		expression = &compiler.Number{Value: p.eat()}

	case lexer.CHAR:
		expression = &compiler.Char{Value: p.eat()}

	case lexer.IDENTIFIER:
		expression = &compiler.Identifier{Value: p.eat()}

	case lexer.TRUE:
		expression = &compiler.True{Value: p.eat()}

	case lexer.FALSE:
		expression = &compiler.False{Value: p.eat()}

	case lexer.OPEN_PAREN:
		p.eat()
		expression = p.parseExpression()
		p.expect(lexer.CLOSE_PAREN, fmt.Sprintf("Syntax Error: expected ) but got %s", logs.TokenToString(p.at())))

	default:
		logs.PrintError(p.at(), fmt.Sprintf("Syntax Error: expected expression but got token %s", logs.TokenToString(p.at())))
		os.Exit(1)
	}

	if p.at().Type == lexer.AS {
		token := p.eat()
		t := p.parseType()

		expression = &compiler.AsExpression{
			Token:      token,
			Type:       t,
			Expression: expression,
		}
	}

	return expression
}
