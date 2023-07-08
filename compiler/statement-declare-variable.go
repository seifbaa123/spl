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
	Expression Node
}

func (d *DeclareVariable) Evaluate(env *Environment) NodeResult {
	if !d.Type.IsValid() {
		logs.PrintError(d.Token, fmt.Sprintf("Type Error: invalid type %s", d.Type.ToString()))
		os.Exit(1)
	}

	address := env.DeclareVariable(d)
	assembly := ""

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
			i.Mov(fmt.Sprintf("[%d]", address), "rax"),
		}, "\n")
	}

	return NodeResult{
		Type: VoidType, Assembly: assembly,
	}
}
