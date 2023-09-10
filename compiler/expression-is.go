package compiler

import (
	"fmt"
	"os"
	"spl/instructions"
	"spl/lexer"
	"spl/logs"
)

type IsExpression struct {
	Token      lexer.Token
	Type       *VariableType
	Expression Node
}

func (i *IsExpression) Evaluate(env *Environment) NodeResult {
	if !i.Type.IsValid() {
		logs.PrintError(i.Token, fmt.Sprintf("Type Error: invalid type %s", i.Type.ToString()))
		os.Exit(1)
	}

	if i.Expression.Evaluate(env).Type == *i.Type {
		return NodeResult{Type: BoolType, Assembly: instructions.Mov("rax", "1")}
	}

	return NodeResult{Type: BoolType, Assembly: instructions.Mov("rax", "0")}
}
