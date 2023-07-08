package compiler

type Node interface {
	Evaluate(env *Environment) NodeResult
}

type NodeResult struct {
	Type     VariableType
	Assembly string
}
