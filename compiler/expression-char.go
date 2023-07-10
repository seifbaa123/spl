package compiler

import (
	"fmt"
	i "spl/instructions"
	"spl/lexer"
)

type Char struct {
	Value lexer.Token
}

func (c *Char) Evaluate(env *Environment) NodeResult {
	return NodeResult{
		Type:     CharType,
		Assembly: i.Mov("rax", fmt.Sprint(int(c.Value.Symbol[0]))),
	}
}
