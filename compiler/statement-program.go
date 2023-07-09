package compiler

import (
	"fmt"
	i "spl/instructions"
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
		Type: VoidType, Assembly: strings.Join(
			[]string{
				i.Sub("rsp", fmt.Sprint(env.address)),
				strings.Join(results, "\n"),
				i.Add("rsp", fmt.Sprint(env.address)),
			},
			"\n",
		),
	}
}
