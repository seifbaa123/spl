package statements

import (
	"spl/node"
)

type Program struct {
	Body []node.Node
}

func (p *Program) Evaluate() node.NodeResult {
	panic("TODO: Program.Evaluate")
}
