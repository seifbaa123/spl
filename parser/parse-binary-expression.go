package parser

import (
	"spl/expressions"
	"spl/lexer"
	"spl/node"
)

func (p *Parser) parseAdding() node.Node {
	left := p.parseMultiplication()

	for p.At().Type == lexer.PLUS || p.At().Type == lexer.MINUS {
		left = &expressions.BinaryExpression{
			Op:    p.Eat(),
			Left:  left,
			Right: p.parseMultiplication(),
		}
	}

	return left
}

func (p *Parser) parseMultiplication() node.Node {
	left := p.ParsePrimary()

	for p.At().Type == lexer.MULTIPLY || p.At().Type == lexer.DIVIDE || p.At().Type == lexer.MODULO {
		left = &expressions.BinaryExpression{
			Op:    p.Eat(),
			Left:  left,
			Right: p.ParsePrimary(),
		}
	}

	return left
}
