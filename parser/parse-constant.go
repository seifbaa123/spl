package parser

import (
	"fmt"
	"os"
	"spl/compiler"
	"spl/lexer"
	"spl/logs"
)

func (p *Parser) parseConstant() *compiler.DeclareVariable {
	constant := p.expect(lexer.CONST, fmt.Sprintf("Syntax Error: expected const keyword but got %s", logs.TokenToString(p.at())))
	name := p.expect(lexer.IDENTIFIER, fmt.Sprintf("Syntax Error: expected constant name but got %s", logs.TokenToString(p.at())))

	p.expect(lexer.COLON, fmt.Sprintf("Syntax Error: expected : after constant name but got %s", logs.TokenToString(p.at())))
	variableType := p.parseType()
	p.expect(lexer.EQUALS, fmt.Sprintf("Syntax Error: expected = after constant type but got %s", logs.TokenToString(p.at())))

	if p.isNewLine() {
		logs.PrintError(
			constant,
			fmt.Sprintf(
				"Syntax Error: expected expression after = token in constant declaration but got %s",
				logs.TokenToString(p.at()),
			),
		)
		os.Exit(1)
	}

	expression := p.parseExpression()

	return &compiler.DeclareVariable{
		Token:      constant,
		Name:       name,
		Type:       variableType,
		IsConstant: true,
		Expression: expression,
	}
}
