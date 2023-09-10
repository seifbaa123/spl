package parser

import (
	"fmt"
	"os"
	"spl/compiler"
	"spl/lexer"
	"spl/logs"
)

func (p *Parser) parseList() *compiler.ListExpression {
	oldIsInParenthesis := p.IsInParenthesis
	p.IsInParenthesis = true

	token := p.expect(lexer.OPEN_BRACKET, fmt.Sprintf("Syntax Error: expected [ but got %s", logs.TokenToString(p.at())))
	var items []compiler.Node

	if p.at().Type != lexer.CLOSE_BRACKET {
		for p.at().Type != lexer.EOF {
			if p.at().Type == lexer.NEW_LINE {
				p.eat()
				continue
			}

			if p.at().Type == lexer.CLOSE_BRACKET {
				break
			}

			item := p.parseExpression()
			if item == nil {
				logs.PrintError(p.at(), "Expected expression in list items!")
				os.Exit(1)
			}
			items = append(items, item)

			if p.at().Type == lexer.CLOSE_BRACKET {
				break
			}

			if p.next().Type != lexer.CLOSE_BRACKET {
				p.expect(lexer.COMMA, fmt.Sprintf("Syntax Error: expected , after list item but got  %s", logs.TokenToString(p.at())))
			}
		}
	}

	p.IsInParenthesis = oldIsInParenthesis
	p.expect(lexer.CLOSE_BRACKET, fmt.Sprintf("Syntax Error: expected ] but got %s", logs.TokenToString(p.at())))

	return &compiler.ListExpression{
		Token: token,
		Items: items,
	}
}
