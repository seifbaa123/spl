package compiler

import (
	"fmt"
	"os"
	i "spl/instructions"
	"spl/lexer"
	"spl/logs"
	"strings"
)

type ReassignVariable struct {
	Op         lexer.Token
	Variable   lexer.Token
	Expression Node
}

func (r *ReassignVariable) Evaluate(env *Environment) NodeResult {
	variable := env.getVariable(r.Variable)
	expression := r.Expression.Evaluate(env)

	if variable.IsConstant {
		logs.PrintError(r.Variable, fmt.Sprintf("Variable Error: can not reassign a constant %s", r.Variable.Symbol))
		os.Exit(1)
	}

	if variable.Type != expression.Type {
		logs.PrintError(
			r.Variable,
			fmt.Sprintf(
				"Type Error: can not assign expression of type %s for variable of type %s",
				expression.Type.ToString(),
				variable.Type.ToString(),
			),
		)
		os.Exit(1)
	}

	return NodeResult{
		Type: VoidType,
		Assembly: strings.Join([]string{
			expression.Assembly,
			i.Mov(fmt.Sprintf("[rbp-%d]", variable.Address), "rax"),
		}, "\n"),
	}
}
