package parser

import (
	"fmt"
	"spl/expressions"
	"spl/lexer"
	"spl/node"
)

func (p *Parser) ParsePrimary() node.Node {
	switch p.At().Type {
	case lexer.NUMBER:
		return &expressions.Number{Value: p.Eat()}
	case lexer.IDENTIFIER:
		return &expressions.Identifier{Value: p.Eat()}
	default:
		fmt.Println(p.At())
		panic("Invalid Token")
	}
}
