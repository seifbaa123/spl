package compiler

import (
	"fmt"
	"os"
	i "spl/instructions"
	"spl/lexer"
	"spl/logs"
	"strings"
)

type AtExpression struct {
	Token      lexer.Token
	Expression Node
	Index      Node
}

func (a *AtExpression) Evaluate(env *Environment) NodeResult {
	expression := a.Expression.Evaluate(env)
	index := a.Index.Evaluate(env)

	// check types
	if !expression.Type.IsValid() {
		logs.PrintError(a.Token, fmt.Sprintf("Type Error: invalid type %s", expression.Type.ToString()))
		os.Exit(1)
	}

	if expression.Type != StrType && expression.Type.Type != "List" {
		logs.PrintError(a.Token, fmt.Sprintf(
			"Type Error: can only use at expression with str or List types but not with %s type",
			expression.Type.ToString()),
		)
		os.Exit(1)
	}

	if !index.Type.IsValid() {
		logs.PrintError(a.Token, fmt.Sprintf("Type Error: invalid type %s", index.Type.ToString()))
		os.Exit(1)
	}

	if index.Type != IntType {
		logs.PrintError(a.Token, fmt.Sprintf(
			"Type Error: the index of at expression must have int type but not %s type",
			index.Type.ToString()),
		)
		os.Exit(1)
	}

	returnType := CharType
	if expression.Type != StrType {
		returnType = *expression.Type.SubType
	}

	// generate assembly
	getAtExpression := i.Movzx("rax", "byte [rbx+8+rax]")
	if expression.Type != StrType {
		getAtExpression = i.Mov("rax", "[rbx+8+rax*8]")
	}

	assembly := strings.Join([]string{
		// check index
		expression.Assembly,
		i.Mov("rax", "[rax]"),
		i.Mov("rbx", "rax"),
		i.Mov("rax", "[rax]"),
		i.Push("rax"),
		index.Assembly,
		i.Push("rax"),
		i.Call("_check_index"),
		i.Add("rsp", "8*2"),

		// get element
		getAtExpression,
	}, "\n")

	return NodeResult{Type: returnType, Assembly: assembly}
}
