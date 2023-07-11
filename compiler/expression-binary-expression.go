package compiler

import (
	"fmt"
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

	checkBinaryExpressionTypes(left, right, b.Op)

	code := []string{
		right.Assembly,
		i.Push("rax"),
		left.Assembly,
		i.Pop("rbx"),
	}

	switch b.Op.Type {
	case lexer.PLUS:
		if left.Type == StrType {
			code = append(code, concatStrings(left, right))
		} else {
			code = append(code, i.Add("rax", "rbx"))
		}

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

	return NodeResult{Type: left.Type, Assembly: strings.Join(code, "\n")}
}

func checkBinaryExpressionTypes(left NodeResult, right NodeResult, op lexer.Token) {
	if right.Type == IntType && left.Type == IntType {
		return
	}

	if right.Type == StrType && left.Type == StrType && op.Type == lexer.PLUS {
		return
	}

	logs.PrintError(
		op,
		fmt.Sprintf(
			"Type Error: can not do operation %s on types %s and %s",
			op.Symbol,
			left.Type.ToString(),
			right.Type.ToString(),
		),
	)
	os.Exit(1)
}

func concatStrings(left NodeResult, right NodeResult) string {
	return strings.Join([]string{
		i.Mov("rax", "[rax]"),
		i.Mov("rbx", "[rbx]"),

		i.Mov("rdx", "rax"),

		// calc new str length
		i.Mov("rcx", "[rdx]"),
		i.Add("rcx", "[rbx]"),

		// allocate new str
		i.Add("rcx", "8"),
		i.Push("rcx"),
		i.Call("_allocate"),
		i.Add("rsp", "8"),
		i.Sub("rcx", "8"),

		// set new str length
		i.Mov("[rax]", "rcx"),

		// copy first str
		i.Mov("rcx", "rdx"),
		i.Add("rcx", "8"),
		i.Push("rcx"),
		i.Mov("rcx", "rax"),
		i.Add("rcx", "8"),
		i.Push("rcx"),
		i.Mov("rcx", "[rdx]"),
		i.Push("rcx"),
		i.Call("_mem_copy"),
		i.Add("rsp", "8*3"),

		// copy second str
		i.Mov("rcx", "rbx"),
		i.Add("rcx", "8"),
		i.Push("rcx"),
		i.Mov("rcx", "rax"),
		i.Add("rcx", "[rdx]"),
		i.Add("rcx", "8"),
		i.Push("rcx"),
		i.Mov("rcx", "[rbx]"),
		i.Push("rcx"),
		i.Call("_mem_copy"),
		i.Add("rsp", "8*3"),

		// allocate str pointer
		i.Mov("rbx", "rax"),
		i.Push("8"),
		i.Call("_allocate"),
		i.Add("rsp", "8"),
		i.Mov("[rax]", "rbx"),
	}, "\n")
}
