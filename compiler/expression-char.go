package compiler

import (
	"fmt"
	"spl/lexer"
)

type Char struct {
	Value lexer.Token
}

func (c *Char) Evaluate(env *Environment) NodeResult {
	return NodeResult{
		Type:     CharType,
		Assembly: fmt.Sprintf("    mov rax, %d", c.Value.Symbol[0]),
	}
}
