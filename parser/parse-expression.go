package parser

import (
	"spl/node"
)

func (p *Parser) parseExpression() node.Node {
	return p.parseAdding()
}
