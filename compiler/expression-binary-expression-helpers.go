package compiler

import (
	"fmt"
	"os"
	i "spl/instructions"
	"spl/lexer"
	"spl/logs"
	"strings"
)

func checkBinaryExpressionTypes(left NodeResult, right NodeResult, op lexer.Token) {
	if right.Type == IntType && left.Type == IntType {
		return
	}

	if right.Type == StrType && left.Type == StrType && op.Type == lexer.PLUS {
		return
	}

	if op.Type == lexer.OR || op.Type == lexer.AND || op.Type == lexer.XOR {
		if right.Type == BoolType && left.Type == BoolType {
			return
		}
	}

	if op.Type == lexer.EQUALS_TO || op.Type == lexer.NOT_EQUALS_TO {
		if (right.Type == StrType && left.Type == StrType) ||
			(right.Type == BoolType && left.Type == BoolType) ||
			(right.Type == CharType && left.Type == CharType) {
			return
		}
	}

	if op.Type == lexer.GREATER || op.Type == lexer.GREATER_OR_EQUALS || op.Type == lexer.LESS || op.Type == lexer.LESS_OR_EQUALS {
		if right.Type == CharType && left.Type == CharType {
			return
		}
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

func handlePlus(left NodeResult, right NodeResult, returnType *VariableType, code *[]string) {
	if left.Type == StrType {
		*returnType = StrType
		*code = append(*code, concatStrings(left, right))
	} else {
		*returnType = IntType
		*code = append(*code, i.Add("rax", "rbx"))
	}
}

func handleMinus(returnType *VariableType, code *[]string) {
	*returnType = IntType
	*code = append(*code, i.Sub("rax", "rbx"))
}

func handleMultiply(returnType *VariableType, code *[]string) {
	*returnType = IntType
	*code = append(*code, i.Mul("rbx"))
}

func handleDivide(returnType *VariableType, code *[]string) {
	*returnType = IntType
	*code = append(*code, strings.Join([]string{
		i.Xor("rdx", "rdx"),
		i.Div("rbx"),
	}, "\n"))
}

func handleModulo(returnType *VariableType, code *[]string) {
	*returnType = IntType
	*code = append(*code, strings.Join([]string{
		i.Xor("rdx", "rdx"),
		i.Div("rbx"),
		i.Mov("rax", "rbx"),
	}, "\n"))
}

func handleOr(returnType *VariableType, code *[]string) {
	*returnType = BoolType
	*code = append(*code, i.Or("rax", "rbx"))
}

func handleAnd(returnType *VariableType, code *[]string) {
	*returnType = BoolType
	*code = append(*code, i.And("rax", "rbx"))
}

func handleXor(returnType *VariableType, code *[]string) {
	*returnType = BoolType
	*code = append(*code, i.Xor("rax", "rbx"))
}

func handleEqualsTo(b *BinaryExpression, left NodeResult, right NodeResult, returnType *VariableType, code *[]string) {
	end := prefixToken(".end", b.Op)
	setTrue := prefixToken(".set_true", b.Op)
	setFalse := prefixToken(".set_false", b.Op)

	*returnType = BoolType
	if left.Type == StrType {
		*code = append(*code, strings.Join([]string{
			i.Push("rax"),
			i.Push("rbx"),
			i.Call("_str_compare"),
			i.Add("rsp", "16"),
		}, "\n"))
	} else {
		*code = append(*code, strings.Join([]string{
			i.Cmp("rax", "rbx"),
			i.Je(setTrue),
			i.Jmp(setFalse),

			setTrue + ":",
			i.Mov("rax", "1"),
			i.Jmp(end),

			setFalse + ":",
			i.Xor("rax", "rax"),

			end + ":",
		}, "\n"))
	}
}

func handleNotEqualsTo(b *BinaryExpression, left NodeResult, right NodeResult, returnType *VariableType, code *[]string) {
	end := prefixToken(".end", b.Op)
	setTrue := prefixToken(".set_true", b.Op)
	setFalse := prefixToken(".set_false", b.Op)

	*returnType = BoolType
	if left.Type == StrType {
		*code = append(*code, strings.Join([]string{
			i.Push("rax"),
			i.Push("rbx"),
			i.Call("_str_compare"),
			i.Add("rsp", "16"),

			i.Cmp("rax", "0"),
			i.Je(setTrue),
			i.Jmp(setFalse),

			setTrue + ":",
			i.Mov("rax", "1"),
			i.Jmp(end),

			setFalse + ":",
			i.Xor("rax", "rax"),

			end + ":",
		}, "\n"))
	} else {
		*code = append(*code, strings.Join([]string{
			i.Cmp("rax", "rbx"),
			i.Jne(setTrue),
			i.Jmp(setFalse),

			setTrue + ":",
			i.Mov("rax", "1"),
			i.Jmp(end),

			setFalse + ":",
			i.Xor("rax", "rax"),

			end + ":",
		}, "\n"))
	}
}

func handleGreater(b *BinaryExpression, left NodeResult, right NodeResult, returnType *VariableType, code *[]string) {
	end := prefixToken(".end", b.Op)
	setTrue := prefixToken(".set_true", b.Op)
	setFalse := prefixToken(".set_false", b.Op)

	*returnType = BoolType
	*code = append(*code, strings.Join([]string{
		i.Cmp("rax", "rbx"),
		i.Jg(setTrue),
		i.Jmp(setFalse),

		setTrue + ":",
		i.Mov("rax", "1"),
		i.Jmp(end),

		setFalse + ":",
		i.Xor("rax", "rax"),

		end + ":",
	}, "\n"))
}

func handleGreaterOrEquals(b *BinaryExpression, left NodeResult, right NodeResult, returnType *VariableType, code *[]string) {
	end := prefixToken(".end", b.Op)
	setTrue := prefixToken(".set_true", b.Op)
	setFalse := prefixToken(".set_false", b.Op)

	*returnType = BoolType
	*code = append(*code, strings.Join([]string{
		i.Cmp("rax", "rbx"),
		i.Jge(setTrue),
		i.Jmp(setFalse),

		setTrue + ":",
		i.Mov("rax", "1"),
		i.Jmp(end),

		setFalse + ":",
		i.Xor("rax", "rax"),

		end + ":",
	}, "\n"))
}

func handleLess(b *BinaryExpression, left NodeResult, right NodeResult, returnType *VariableType, code *[]string) {
	end := prefixToken(".end", b.Op)
	setTrue := prefixToken(".set_true", b.Op)
	setFalse := prefixToken(".set_false", b.Op)

	*returnType = BoolType
	*code = append(*code, strings.Join([]string{
		i.Cmp("rax", "rbx"),
		i.Jl(setTrue),
		i.Jmp(setFalse),

		setTrue + ":",
		i.Mov("rax", "1"),
		i.Jmp(end),

		setFalse + ":",
		i.Xor("rax", "rax"),

		end + ":",
	}, "\n"))
}

func handleLessOrEquals(b *BinaryExpression, left NodeResult, right NodeResult, returnType *VariableType, code *[]string) {
	end := prefixToken(".end", b.Op)
	setTrue := prefixToken(".set_true", b.Op)
	setFalse := prefixToken(".set_false", b.Op)

	*returnType = BoolType
	*code = append(*code, strings.Join([]string{
		i.Cmp("rax", "rbx"),
		i.Jle(setTrue),
		i.Jmp(setFalse),

		setTrue + ":",
		i.Mov("rax", "1"),
		i.Jmp(end),

		setFalse + ":",
		i.Xor("rax", "rax"),

		end + ":",
	}, "\n"))
}
