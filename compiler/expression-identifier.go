package compiler

import (
	"spl/lexer"
)

type Identifier struct {
	Value lexer.Token
}

func (i *Identifier) Evaluate(env *Environment) NodeResult {
	panic("TODO: Identifier.Evaluate")
}
