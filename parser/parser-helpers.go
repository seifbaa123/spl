package parser

import "spl/lexer"

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
