package compiler

import (
	i "spl/instructions"
	"spl/lexer"
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

	var returnType VariableType

	switch b.Op.Type {
	// arithmetic operations
	case lexer.PLUS:
		handlePlus(left, right, &returnType, &code)
	case lexer.MINUS:
		handleMinus(&returnType, &code)
	case lexer.MULTIPLY:
		handleMultiply(&returnType, &code)
	case lexer.DIVIDE:
		handleDivide(&returnType, &code)
	case lexer.MODULO:
		handleModulo(&returnType, &code)

	// logical operations
	case lexer.OR:
		handleOr(&returnType, &code)
	case lexer.AND:
		handleAnd(&returnType, &code)
	case lexer.XOR:
		handleXor(&returnType, &code)

	case lexer.EQUALS_TO:
		handleEqualsTo(b, left, right, &returnType, &code)
	case lexer.NOT_EQUALS_TO:
		handleNotEqualsTo(b, left, right, &returnType, &code)

	case lexer.GREATER:
		handleGreater(b, left, right, &returnType, &code)
	case lexer.GREATER_OR_EQUALS:
		handleGreaterOrEquals(b, left, right, &returnType, &code)
	case lexer.LESS:
		handleLess(b, left, right, &returnType, &code)
	case lexer.LESS_OR_EQUALS:
		handleLessOrEquals(b, left, right, &returnType, &code)
	}

	return NodeResult{Type: returnType, Assembly: strings.Join(code, "\n")}
}
