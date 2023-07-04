package node

type Node interface {
	Evaluate() NodeResult
}

type NodeResult struct {
	Assembly string
}
