package compiler

import (
	"spl/lexer"
)

type True struct {
	Value lexer.Token
}

type False struct {
	Value lexer.Token
}

func (t *True) Evaluate(env *Environment) NodeResult {
	return NodeResult{
		Type:     BoolType,
		Assembly: "    mov rax, 1",
	}
}

func (f *False) Evaluate(env *Environment) NodeResult {
	return NodeResult{
		Type:     BoolType,
		Assembly: "    mov rax, 0",
	}
}
