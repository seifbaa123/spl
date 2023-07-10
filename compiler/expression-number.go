package compiler

import (
	"fmt"
	"spl/lexer"
)

type Int struct {
	Value lexer.Token
}

func (n *Int) Evaluate(env *Environment) NodeResult {
	return NodeResult{
		Type:     IntType,
		Assembly: fmt.Sprintf("    mov rax, %s", n.Value.Symbol),
	}
}
