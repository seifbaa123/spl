package expressions

import (
	"spl/instructions"
	"spl/lexer"
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

	code := []string{
		right.Assembly,
		instructions.Push("rax"),
		left.Assembly,
		instructions.Pop("rbx"),
	}

	switch b.Op.Type {
	case lexer.PLUS:
		code = append(code, instructions.Add("rax", "rbx"))
	case lexer.MINUS:
		code = append(code, instructions.Sub("rax", "rbx"))
	case lexer.MULTIPLY:
		code = append(code, instructions.Mul("rbx"))
	case lexer.DIVIDE:
		code = append(code, strings.Join([]string{
			instructions.Xor("rdx", "rdx"),
			instructions.Div("rbx"),
		}, "\n"))
	case lexer.MODULO:
		code = append(code, strings.Join([]string{
			instructions.Xor("rdx", "rdx"),
			instructions.Div("rbx"),
			instructions.Mov("rax", "rbx"),
		}, "\n"))
	}

	return node.NodeResult{Assembly: strings.Join(code, "\n")}
}
