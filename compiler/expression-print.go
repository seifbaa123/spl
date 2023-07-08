package compiler

import (
	"os"
	i "spl/instructions"
	"spl/lexer"
	"spl/logs"
	"strings"
)

type Print struct {
	Token      lexer.Token
	Expression Node
}

func (p *Print) Evaluate(env *Environment) NodeResult {
	expression := p.Expression.Evaluate(env)

	if expression.Type == VoidType {
		logs.PrintError(p.Token, "Type Error: can not print expression of type void")
		os.Exit(1)
	}

	return NodeResult{
		Type: VoidType,
		Assembly: strings.Join([]string{
			expression.Assembly,
			i.Push("rax"),
			i.Call(getPrintFunctionName(expression.Type)),
			i.Add("rsp", "8"),
		}, "\n"),
	}
}

func getPrintFunctionName(t VariableType) string {
	switch t {
	case CharType:
		return "_print_char"
	case BoolType:
		return "_print_bool"
	}

	return "_print_int"
}
