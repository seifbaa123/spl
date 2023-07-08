package compiler

import (
	"os"
	i "spl/instructions"
	"spl/lexer"
	"spl/logs"
	"strings"
)

type BinaryExpression struct {
	Op    lexer.Token
	Left  Node
	Right Node
}

func (b *BinaryExpression) Evaluate(env *Environment) NodeResult {
	left := b.Left.Evaluate(env)
	right := b.Right.Evaluate(env)

	if right.Type != IntType || left.Type != IntType {
		logs.PrintError(b.Op, "Type Error: binary operation can only implied on numerical types")
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

	return NodeResult{Type: IntType, Assembly: strings.Join(code, "\n")}
}
