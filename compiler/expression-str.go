package compiler

import (
	"fmt"
	i "spl/instructions"
	"spl/lexer"
	"strings"
)

type Str struct {
	Value lexer.Token
}

func (s *Str) Evaluate(env *Environment) NodeResult {
	str := fmt.Sprintf("str%d", addString(s.Value.Symbol))

	return NodeResult{
		Type: StrType,
		Assembly: strings.Join([]string{
			// allocate string bytes
			i.Push(fmt.Sprintf("8+%d+1", len(s.Value.Symbol))), // 8 bytes for length + string bytes + null byte
			i.Call("_allocate"),
			i.Add("rsp", "8"),

			// set length
			i.Mov("rbx", fmt.Sprint(len(s.Value.Symbol))),
			i.Mov("[rax]", "rbx"),

			// copy string
			i.Push(str),
			i.Add("rax", "8"),
			i.Push("rax"),
			i.Sub("rax", "8"),
			i.Push("rbx"),
			i.Call("_mem_copy"),
			i.Add("rsp", "8*3"),
			i.Mov(fmt.Sprintf("[rax+8+%d]", len(s.Value.Symbol)), "byte 0"),

			// set string pointer
			i.Mov("rbx", "rax"),
			i.Push("8"),
			i.Call("_allocate"),
			i.Add("rsp", "8"),
			i.Mov("[rax]", "rbx"),
		}, "\n"),
	}
}
