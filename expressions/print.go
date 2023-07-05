package expressions

import (
	"os"
	i "spl/instructions"
	"spl/lexer"
	"spl/logs"
	"spl/node"
	"strings"
)

type Print struct {
	Token      lexer.Token
	Expression node.Node
}

func (p *Print) Evaluate() node.NodeResult {
	expression := p.Expression.Evaluate()

	if expression.Type == node.Void {
		logs.PrintError(p.Token, "Can not print expression of type void")
		os.Exit(1)
	}

	return node.NodeResult{
		Type: node.Void,
		Assembly: strings.Join([]string{
			expression.Assembly,
			i.Push("rax"),
			i.Call(getPrintFunctionName(expression.Type)),
			i.Add("rsp", "8"),
		}, "\n"),
	}
}

func getPrintFunctionName(t node.VariableType) string {
	switch t {
	case node.Char:
		return "_print_char"
	case node.Bool:
		return "_print_bool"
	}

	return "_print_int"
}
