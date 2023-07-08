package parser

import (
	"fmt"
	"os"
	"spl/compiler"
	"spl/lexer"
	"spl/logs"
)

func (p *Parser) ParsePrimary() compiler.Node {
	switch p.at().Type {
	case lexer.NUMBER:
		return &compiler.Number{Value: p.eat()}

	case lexer.CHAR:
		return &compiler.Char{Value: p.eat()}

	case lexer.IDENTIFIER:
		return &compiler.Identifier{Value: p.eat()}

	case lexer.TRUE:
		return &compiler.True{Value: p.eat()}

	case lexer.FALSE:
		return &compiler.False{Value: p.eat()}

	case lexer.OPEN_PAREN:
		p.eat()
		expression := p.parseExpression()
		p.expect(lexer.CLOSE_PAREN, fmt.Sprintf("Syntax Error: expected ) but got %s", p.at().Symbol))
		return expression

	default:
		logs.PrintError(p.at(), fmt.Sprintf("Syntax Error: unexpected Token %s", logs.TokenToString(p.at())))
		os.Exit(1)
	}

	return nil
}
