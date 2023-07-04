package parser

import (
	"fmt"
	"os"
	"spl/lexer"
	"spl/logs"
)

func (p *Parser) At() lexer.Token {
	if p.Index >= uint(len(p.Tokens)) {
		return lexer.Token{Type: lexer.EOF}
	}

	return p.Tokens[p.Index]
}

func (p *Parser) Eat() lexer.Token {
	token := p.At()
	p.Index++

	return token
}

func (p *Parser) Expect(tokenType lexer.TokenType, message string) lexer.Token {
	if p.At().Type != tokenType {
		logs.PrintError(p.At(), message)
		os.Exit(1)
	}

	return p.Eat()
}

func (p *Parser) ExpectNewLine() {
	if p.At().Type != lexer.SEMI_COLON && p.At().Type != lexer.END_OF_LINE && p.At().Type != lexer.EOF {
		logs.PrintError(p.At(), fmt.Sprintf("Expected new line or ; but got %s", logs.TokenToString(p.At())))
		os.Exit(1)
	}

	p.Eat()
}
