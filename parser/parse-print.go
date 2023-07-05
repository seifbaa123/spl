package parser

import (
	"fmt"
	"spl/expressions"
	"spl/lexer"
	"spl/logs"
)

func (p *Parser) parsePrint() *expressions.Print {
	token := p.Expect(
		lexer.PRINT,
		fmt.Sprintf("Expected token print but got %s", logs.TokenToString(p.At())),
	)

	p.Expect(lexer.OPEN_PAREN, fmt.Sprintf("Syntax Error: expected ( but got %s", logs.TokenToString(p.At())))
	expression := p.parseExpression()
	p.Expect(lexer.CLOSE_PAREN, fmt.Sprintf("Syntax Error: expected ) but got %s", logs.TokenToString(p.At())))

	return &expressions.Print{
		Token:      token,
		Expression: expression,
	}
}
