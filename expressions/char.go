package expressions

import (
	"fmt"
	"spl/lexer"
	"spl/node"
)

type Char struct {
	Value lexer.Token
}

func (c *Char) Evaluate() node.NodeResult {
	return node.NodeResult{
		Type:     node.Char,
		Assembly: fmt.Sprintf("    mov rax, %d", c.Value.Symbol[0]),
	}
}
