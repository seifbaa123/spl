package expressions

import (
	"spl/lexer"
	"spl/node"
)

type Number struct {
	Value lexer.Token
}

func (n *Number) Evaluate() node.NodeResult {
	panic("TODO: Number.Evaluate")
}
