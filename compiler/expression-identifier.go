package compiler

import (
	"fmt"
	"spl/instructions"
	"spl/lexer"
)

type Identifier struct {
	Value lexer.Token
}

func (i *Identifier) Evaluate(env *Environment) NodeResult {
	variable := env.getVariable(i.Value)

	return NodeResult{
		Type:     variable.Type,
		Assembly: instructions.Mov("rax", fmt.Sprintf("[rbp-%d]", variable.Address)),
	}
}
