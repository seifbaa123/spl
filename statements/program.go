package statements

import (
	"spl/node"
	"strings"
)

type Program struct {
	Body []node.Node
}

func (p *Program) Evaluate() node.NodeResult {
	var results []string

	for _, node := range p.Body {
		results = append(results, node.Evaluate().Assembly)
	}

	return node.NodeResult{
		Assembly: strings.Join(results, "\n"),
	}
}
