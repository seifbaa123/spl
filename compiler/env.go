package compiler

import (
	"fmt"
	"os"
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

func (e *Environment) DeclareVariable(variable *DeclareVariable) uint {
	_, isExist := e.Variables[variable.Name.Symbol]

	if isExist {
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
