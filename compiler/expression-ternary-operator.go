package compiler

import (
	"fmt"
	"os"
	i "spl/instructions"
	"spl/lexer"
	"spl/logs"
	"strings"
)

type TernaryOperator struct {
	Token          lexer.Token
	Condition      Node
	IfExpression   Node
	ElseExpression Node
}

func (t *TernaryOperator) Evaluate(env *Environment) NodeResult {
	condition := t.Condition.Evaluate(env)
	if condition.Type != BoolType {
		logs.PrintError(
			t.Token,
			fmt.Sprintf("TypeError: condition expression must be of type boolean but its of type %s", condition.Type.ToString()),
		)
		os.Exit(1)
	}

	ifExpression := t.IfExpression.Evaluate(env)
	elseExpression := t.ElseExpression.Evaluate(env)

	if !ifExpression.Type.Compare(elseExpression.Type) {
		logs.PrintError(
			t.Token,
			"TypeError: the if and else expression in ternary operator expression must have the same type",
		)
		os.Exit(1)
	}

	elseLabel := prefixToken(".else", t.Token)
	endLabel := prefixToken(".end", t.Token)

	return NodeResult{
		Type: ifExpression.Type,
		Assembly: strings.Join([]string{
			condition.Assembly,

			i.Cmp("rax", "0"),
			i.Je(elseLabel),

			ifExpression.Assembly,
			i.Jmp(endLabel),

			elseLabel + ":",
			elseExpression.Assembly,

			endLabel + ":",
		}, "\n"),
	}
}
