package expressions

import (
	"spl/lexer"
	"spl/node"
)

type True struct {
	Value lexer.Token
}

type False struct {
	Value lexer.Token
}

func (t *True) Evaluate() node.NodeResult {
	return node.NodeResult{
		Type:     node.Bool,
		Assembly: "    mov rax, 1",
	}
}

func (f *False) Evaluate() node.NodeResult {
	return node.NodeResult{
		Type:     node.Bool,
		Assembly: "    mov rax, 0",
	}
}
