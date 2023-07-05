package expressions

import (
	"os"
	i "spl/instructions"
	"spl/lexer"
	"spl/logs"
	"spl/node"
	"strings"
)

type BinaryExpression struct {
	Op    lexer.Token
	Left  node.Node
	Right node.Node
}

func (b *BinaryExpression) Evaluate() node.NodeResult {
	left := b.Left.Evaluate()
	right := b.Right.Evaluate()

	if right.Type != node.Int || left.Type != node.Int {
		logs.PrintError(b.Op, "Binary operation can only implied on numerical types")
		os.Exit(1)
	}

	code := []string{
		right.Assembly,
		i.Push("rax"),
		left.Assembly,
		i.Pop("rbx"),
	}

	switch b.Op.Type {
	case lexer.PLUS:
		code = append(code, i.Add("rax", "rbx"))
	case lexer.MINUS:
		code = append(code, i.Sub("rax", "rbx"))
	case lexer.MULTIPLY:
		code = append(code, i.Mul("rbx"))
	case lexer.DIVIDE:
		code = append(code, strings.Join([]string{
			i.Xor("rdx", "rdx"),
			i.Div("rbx"),
		}, "\n"))
	case lexer.MODULO:
		code = append(code, strings.Join([]string{
			i.Xor("rdx", "rdx"),
			i.Div("rbx"),
			i.Mov("rax", "rbx"),
		}, "\n"))
	}

	return node.NodeResult{Type: node.Int, Assembly: strings.Join(code, "\n")}
}
