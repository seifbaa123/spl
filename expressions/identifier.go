package expressions

import (
	"spl/lexer"
	"spl/node"
)

type Identifier struct {
	Value lexer.Token
}

func (i *Identifier) Evaluate() node.NodeResult {
	panic("TODO: Identifier.Evaluate")
}
