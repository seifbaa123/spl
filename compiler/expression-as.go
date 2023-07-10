package compiler

import (
	"fmt"
	"os"
	"spl/lexer"
	"spl/logs"
)

type AsExpression struct {
	Token      lexer.Token
	Type       VariableType
	Expression Node
}

func (a *AsExpression) Evaluate(env *Environment) NodeResult {
	if !a.Type.IsValid() {
		logs.PrintError(a.Token, fmt.Sprintf("Type Error: invalid type %s", a.Type.ToString()))
		os.Exit(1)
	}

	return NodeResult{Type: a.Type, Assembly: a.Expression.Evaluate(env).Assembly}
}
