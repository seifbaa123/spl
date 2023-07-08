package compiler

import (
	"fmt"
	"spl/lexer"
)

type Number struct {
	Value lexer.Token
}

func (n *Number) Evaluate(env *Environment) NodeResult {
	return NodeResult{
		Type:     IntType,
		Assembly: fmt.Sprintf("    mov rax, %s", n.Value.Symbol),
	}
}
