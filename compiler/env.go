package compiler

import (
	"fmt"
	"os"
	"spl/lexer"
	"spl/logs"
)

type Variable struct {
	Address uint
	Type    VariableType
}

type Environment struct {
	Variables map[string]Variable
	address   uint
}

func NewEnvironment() *Environment {
	return &Environment{
		Variables: make(map[string]Variable),
		address:   0,
	}
}

func (e *Environment) declareVariable(variable *DeclareVariable) uint {
	if e.hasVariable(variable.Name.Symbol) {
		logs.PrintError(variable.Token, fmt.Sprintf("Variable Error: variable %s is already declared", variable.Name.Symbol))
		os.Exit(1)
	}

	address := e.address
	e.address += 8

	e.Variables[variable.Name.Symbol] = Variable{
		Address: address,
		Type:    variable.Type,
	}

	return address
}

func (e *Environment) getVariable(variable lexer.Token) Variable {
	if !e.hasVariable(variable.Symbol) {
		logs.PrintError(variable, fmt.Sprintf("Variable Error: variable %s is not declared", variable.Symbol))
		os.Exit(1)
	}

	v := e.Variables[variable.Symbol]
	return v
}

func (e *Environment) hasVariable(variable string) bool {
	_, isExist := e.Variables[variable]
	return isExist
}
