package parser

import (
	"fmt"
	"os"
	"spl/lexer"
	"spl/logs"
)

func (p *Parser) at() lexer.Token {
	if p.Index >= uint(len(p.Tokens)) {
		return lexer.Token{Type: lexer.EOF}
	}

	return p.Tokens[p.Index]
}

func (p *Parser) eat() lexer.Token {
	token := p.at()
	p.Index++

	return token
}

func (p *Parser) expect(tokenType lexer.TokenType, message string) lexer.Token {
	if p.at().Type != tokenType {
		logs.PrintError(p.at(), message)
		os.Exit(1)
	}

	return p.eat()
}

func (p *Parser) expectNewLine() {
	if !p.isNewLine() {
		logs.PrintError(p.at(), fmt.Sprintf("Syntax Error: expected new line or ; but got %s", logs.TokenToString(p.at())))
		os.Exit(1)
	}

	p.eat()
}

func (p *Parser) isNewLine() bool {
	return p.at().Type == lexer.SEMI_COLON || p.at().Type == lexer.END_OF_LINE || p.at().Type == lexer.EOF
}
