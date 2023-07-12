package compiler

import (
	i "spl/instructions"
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
		Assembly: i.Mov("rax", "1"),
	}
}

func (f *False) Evaluate(env *Environment) NodeResult {
	return NodeResult{
		Type:     BoolType,
		Assembly: i.Xor("rax", "rax"),
	}
}
