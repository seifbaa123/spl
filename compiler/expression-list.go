package compiler

import (
	"fmt"
	"os"
	i "spl/instructions"
	"spl/lexer"
	"spl/logs"
	"strings"
)

type ListExpression struct {
	Token lexer.Token
	Items []Node
}

func (l *ListExpression) Evaluate(env *Environment) NodeResult {
	var assembly string
	listType := VariableType{Type: "List"}

	if len(l.Items) == 0 {
		assembly = getEmptyListTypeAndAssembly(&listType)
	} else {
		assembly = getListTypeAndAssembly(l, &listType, env)
	}

	return NodeResult{Type: listType, Assembly: assembly}
}

func getListTypeAndAssembly(l *ListExpression, listType *VariableType, env *Environment) string {
	firstItemType := l.Items[0].Evaluate(env).Type
	listType.SubType = &firstItemType

	var items []NodeResult

	for _, i := range l.Items {
		item := i.Evaluate(env)
		items = append(items, item)

		if !item.Type.Compare(firstItemType) {
			logs.PrintError(
				l.Token,
				fmt.Sprintf(
					"TypeError: List can not contain items from type %s and type %s",
					firstItemType.ToString(), item.Type.ToString(),
				),
			)
			os.Exit(1)
		}
	}

	var itemsAssembly []string
	for index, item := range items {
		itemsAssembly = append(itemsAssembly, strings.Join([]string{
			i.Push("rax"),
			item.Assembly,
			i.Pop("rbx"),
			i.Mov(fmt.Sprintf("[rbx+8+%d]", index), "rax"),
			i.Mov("rax", "rbx"),
		}, "\n"))
	}

	return strings.Join([]string{
		// allocate empty list
		i.Push(fmt.Sprintf("8+%d*8", len(l.Items))), // 8 bytes for length + list bytes,
		i.Call("_allocate"),
		i.Add("rsp", "8"),

		// set length
		i.Mov("rbx", fmt.Sprintf("%d", len(l.Items))),
		i.Mov("[rax]", "rbx"),

		// set items
		strings.Join(itemsAssembly, "\n"),

		// set pointer
		i.Mov("rbx", "rax"),
		i.Push("8"),
		i.Call("_allocate"),
		i.Add("rsp", "8"),
		i.Mov("[rax]", "rbx"),
	}, "\n")
}

func getEmptyListTypeAndAssembly(listType *VariableType) string {
	listType.IsEmptyList = true
	return strings.Join([]string{
		// allocate empty list
		i.Push("8"),
		i.Call("_allocate"),
		i.Add("rsp", "8"),

		// set length
		i.Mov("rbx", "0"),
		i.Mov("[rax]", "rbx"),

		// set pointer
		i.Mov("rbx", "rax"),
		i.Push("8"),
		i.Call("_allocate"),
		i.Add("rsp", "8"),
		i.Mov("[rax]", "rbx"),
	}, "\n")
}
