package parser

import (
	"fmt"
	"os"
	"spl/expressions"
	"spl/lexer"
	"spl/logs"
	"spl/node"
)

func (p *Parser) ParsePrimary() node.Node {
	switch p.At().Type {
	case lexer.NUMBER:
		return &expressions.Number{Value: p.Eat()}
	case lexer.CHAR:
		return &expressions.Char{Value: p.Eat()}
	case lexer.IDENTIFIER:
		return &expressions.Identifier{Value: p.Eat()}
	case lexer.TRUE:
		return &expressions.True{Value: p.Eat()}
	case lexer.FALSE:
		return &expressions.False{Value: p.Eat()}
	default:
		logs.PrintError(p.At(), fmt.Sprintf("Unexpected Token %s", logs.TokenToString(p.At())))
		os.Exit(1)
	}

	return nil
}
