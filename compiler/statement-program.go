package compiler

import (
	"strings"
)

type Program struct {
	Body []Node
}

func (p *Program) Evaluate(env *Environment) NodeResult {
	var results []string

	for _, node := range p.Body {
		results = append(results, node.Evaluate(env).Assembly)
	}

	return NodeResult{
		Type: VoidType, Assembly: strings.Join(results, "\n"),
	}
}
