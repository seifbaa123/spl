package compiler

import (
	"fmt"
	"os"
	"spl/instructions"
	"spl/lexer"
	"spl/logs"
	"strings"
)

type Increment struct {
	Operator   lexer.Token
	Identifier lexer.Token
	IsAfter    bool
}

func (i *Increment) Evaluate(env *Environment) NodeResult {
	variable := env.getVariable(i.Identifier)

	if variable.Type != IntType {
		logs.PrintError(
			i.Operator,
			fmt.Sprintf(
				"TypeError: can not use operator %s with type %s, it can be used only with int type",
				i.Operator.Symbol, variable.Type.ToString(),
			),
		)
		os.Exit(1)
	}

	address := fmt.Sprintf("[rbp-%d]", variable.Address)
	instruction := instructions.Add
	if i.Operator.Type == lexer.MINUS_MINUS {
		instruction = instructions.Sub
	}

	if i.IsAfter {
		return NodeResult{
			Type: IntType,
			Assembly: strings.Join([]string{
				instructions.Mov("rax", address),
				instructions.Mov("rbx", "rax"),
				instruction("rbx", "1"),
				instructions.Mov(address, "rbx"),
			}, "\n"),
		}
	}

	return NodeResult{
		Type: IntType,
		Assembly: strings.Join([]string{
			instructions.Mov("rax", address),
			instruction("rax", "1"),
			instructions.Mov(address, "rax"),
		}, "\n"),
	}
}
