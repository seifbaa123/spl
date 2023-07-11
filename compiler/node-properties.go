package compiler

import (
	i "spl/instructions"
	"strings"
)

type Property struct {
	Type     VariableType
	Assembly string
}

type PropertiesMap map[string]Property

var propertiesList = map[string]PropertiesMap{
	"str": {
		"length": {
			Type: IntType,
			Assembly: strings.Join([]string{
				i.Mov("rax", "[rax]"),
				i.Mov("rax", "[rax]"),
			}, "\n"),
		},
	},
}
