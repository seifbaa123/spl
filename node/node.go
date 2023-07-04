package node

type Node interface {
	Evaluate() NodeResult
}

type NodeResult struct {
	Type     VariableType
	Assembly string
}
