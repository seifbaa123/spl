package compiler

import (
	"os"
	i "spl/instructions"
	"spl/lexer"
	"spl/logs"
	"strings"
)

type NotExpression struct {
	Token      lexer.Token
	Expression Node
}

func (n *NotExpression) Evaluate(env *Environment) NodeResult {
	expression := n.Expression.Evaluate(env)

	if expression.Type != BoolType {
		logs.PrintError(n.Token, "Type Error: can not use the not operator with non boolean types")
		os.Exit(1)
	}

	endLabel := prefixToken(".end", n.Token)
	trueLabel := prefixToken(".set_true", n.Token)
	falseLabel := prefixToken(".set_false", n.Token)

	return NodeResult{
		Type: BoolType,
		Assembly: strings.Join([]string{
			expression.Assembly,
			i.Cmp("rax", "0"),
			i.Je(trueLabel),
			i.Jmp(falseLabel),

			trueLabel + ":",
			i.Mov("rax", "1"),
			i.Jmp(endLabel),

			falseLabel + ":",
			i.Xor("rax", "rax"),

			endLabel + ":",
		}, "\n"),
	}
}
