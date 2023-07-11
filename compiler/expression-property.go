package compiler

import (
	"fmt"
	"os"
	"spl/lexer"
	"spl/logs"
	"strings"
)

type PropertyExpression struct {
	Property   lexer.Token
	Expression Node
}

func (p *PropertyExpression) Evaluate(env *Environment) NodeResult {
	expression := p.Expression.Evaluate(env)
	properties, typeExist := propertiesList[expression.Type.Type]

	if !typeExist {
		logs.PrintError(
			p.Property,
			fmt.Sprintf(
				"Type Error: expression of type %s does not contain any properties",
				expression.Type.ToString(),
			),
		)
		os.Exit(1)
	}

	property, propertyExist := properties[p.Property.Symbol]

	if !propertyExist {
		logs.PrintError(
			p.Property,
			fmt.Sprintf(
				"Type Error: expression of type %s does not contain property %s",
				expression.Type.ToString(),
				p.Property.Symbol,
			),
		)
		os.Exit(1)
	}

	return NodeResult{
		Type: property.Type,
		Assembly: strings.Join([]string{
			expression.Assembly,
			property.Assembly,
		}, "\n"),
	}
}
