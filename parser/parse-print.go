package parser

import (
	"fmt"
	"spl/compiler"
	"spl/lexer"
	"spl/logs"
)

func (p *Parser) parsePrint() *compiler.Print {
	token := p.expect(
		lexer.PRINT,
		fmt.Sprintf("expected token print but got %s", logs.TokenToString(p.at())),
	)

	p.expect(lexer.OPEN_PAREN, fmt.Sprintf("Syntax Error: expected ( but got %s", logs.TokenToString(p.at())))
	expression := p.parseExpression()
	p.expect(lexer.CLOSE_PAREN, fmt.Sprintf("Syntax Error: expected ) but got %s", logs.TokenToString(p.at())))

	return &compiler.Print{
		Token:      token,
		Expression: expression,
	}
}
