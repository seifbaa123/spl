package compiler

import (
	"fmt"
	i "spl/instructions"
	"spl/lexer"
)

type Int struct {
	Value lexer.Token
}

func (n *Int) Evaluate(env *Environment) NodeResult {
	return NodeResult{
		Type:     IntType,
		Assembly: i.Mov("rax", fmt.Sprint(n.Value.Symbol)),
	}
}
