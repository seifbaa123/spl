package parser

import (
	"fmt"
	"spl/compiler"
	"spl/lexer"
	"spl/logs"
)

func (p *Parser) parseVariable() *compiler.DeclareVariable {
	let := p.expect(lexer.LET, fmt.Sprintf("Syntax Error: expected let keyword but got %s", logs.TokenToString(p.at())))
	name := p.expect(lexer.IDENTIFIER, fmt.Sprintf("Syntax Error: expected variable name but got %s", logs.TokenToString(p.at())))

	p.expect(lexer.COLON, fmt.Sprintf("Syntax Error: expected : after variable name but got %s", logs.TokenToString(p.at())))
	variableType := p.parseType()
	p.expect(lexer.EQUALS, fmt.Sprintf("Syntax Error: expected = after variable type but got %s", logs.TokenToString(p.at())))

	var expression compiler.Node
	if !p.isNewLine() {
		expression = p.parseExpression()
	}

	return &compiler.DeclareVariable{
		Token:      let,
		Name:       name,
		Type:       variableType,
		IsConstant: false,
		Expression: expression,
	}
}
