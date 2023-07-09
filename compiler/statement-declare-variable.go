package compiler

import (
	"fmt"
	"os"
	i "spl/instructions"
	"spl/lexer"
	"spl/logs"
	"strings"
)

type DeclareVariable struct {
	Token      lexer.Token
	Name       lexer.Token
	Type       VariableType
	IsConstant bool
	Expression Node
}

func (d *DeclareVariable) Evaluate(env *Environment) NodeResult {
	if !d.Type.IsValid() {
		logs.PrintError(d.Token, fmt.Sprintf("Type Error: invalid type %s", d.Type.ToString()))
		os.Exit(1)
	}

	var assembly string
	address := env.declareVariable(d)

	if d.Expression != nil {
		expression := d.Expression.Evaluate(env)

		if expression.Type != d.Type {
			logs.PrintError(
				d.Token,
				fmt.Sprintf("Type Error: can not assign expression of type %s to variable of type %s",
					expression.Type.ToString(),
					d.Type.ToString()),
			)
		}

		assembly = strings.Join([]string{
			expression.Assembly,
			i.Mov(fmt.Sprintf("[rbp-%d]", address), "rax"),
		}, "\n")
	} else {
		assembly = i.Mov(fmt.Sprintf("[rbp-%d]", address), "0")
	}

	return NodeResult{
		Type: VoidType, Assembly: assembly,
	}
}
