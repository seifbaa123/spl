package parser

import (
	"fmt"
	"os"
	"spl/compiler"
	"spl/lexer"
	"spl/logs"
)

func (p *Parser) parseAt(expression compiler.Node) *compiler.AtExpression {
	token := p.expect(lexer.OPEN_BRACKET, fmt.Sprintf("SyntaxError: expected [ but got %s", logs.TokenToString(p.at())))

	index := p.parseExpression()
	if index == nil {
		logs.PrintError(p.at(), "Expected expression after [")
		os.Exit(0)
	}

	p.expect(lexer.CLOSE_BRACKET, fmt.Sprintf("SyntaxError: expected ] but got %s", logs.TokenToString(p.at())))

	return &compiler.AtExpression{
		Token:      token,
		Expression: expression,
		Index:      index,
	}
}
