package parser

import (
	"spl/compiler"
	"spl/lexer"
)

func (p *Parser) parseLogical() compiler.Node {
	left := p.parseAdding()

	for p.at().Type == lexer.OR || p.at().Type == lexer.AND || p.at().Type == lexer.XOR {
		left = &compiler.BinaryExpression{
			Op:    p.eat(),
			Left:  left,
			Right: p.parseAdding(),
		}
	}

	return left
}

func (p *Parser) parseAdding() compiler.Node {
	left := p.parseMultiplication()

	for p.at().Type == lexer.PLUS || p.at().Type == lexer.MINUS {
		left = &compiler.BinaryExpression{
			Op:    p.eat(),
			Left:  left,
			Right: p.parseMultiplication(),
		}
	}

	return left
}

func (p *Parser) parseMultiplication() compiler.Node {
	left := p.ParsePrimary()

	for p.at().Type == lexer.MULTIPLY || p.at().Type == lexer.DIVIDE || p.at().Type == lexer.MODULO {
		left = &compiler.BinaryExpression{
			Op:    p.eat(),
			Left:  left,
			Right: p.ParsePrimary(),
		}
	}

	return left
}
