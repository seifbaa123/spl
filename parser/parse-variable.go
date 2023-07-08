package parser

import (
	"fmt"
	"spl/compiler"
	"spl/lexer"
)

func (p *Parser) parseVariable() *compiler.DeclareVariable {
	let := p.expect(lexer.LET, fmt.Sprintf("Syntax Error: expected let keyword but got %s", p.at().Symbol))
	name := p.expect(lexer.IDENTIFIER, fmt.Sprintf("Syntax Error: expected variable name but got %s", p.at().Symbol))

	p.expect(lexer.COLON, fmt.Sprintf("Syntax Error: expected : after variable name but got %s", p.at().Symbol))
	variableType := p.parseType()
	p.expect(lexer.EQUALS, fmt.Sprintf("Syntax Error: expected = after variable type but got %s", p.at().Symbol))

	var expression compiler.Node
	if !p.isNewLine() {
		expression = p.parseExpression()
	}

	return &compiler.DeclareVariable{
		Token:      let,
		Name:       name,
		Type:       variableType,
		Expression: expression,
	}
}
