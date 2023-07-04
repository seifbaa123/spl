package expressions

import (
	"spl/lexer"
	"spl/node"
)

type BinaryExpression struct {
	Op    lexer.Token
	Left  node.Node
	Right node.Node
}

func (b *BinaryExpression) Evaluate() node.NodeResult {
	panic("TODO: BinaryExpression.Evaluate")
}
