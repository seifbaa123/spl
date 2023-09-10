package compiler

import "fmt"

var IntType = VariableType{Type: "int"}
var StrType = VariableType{Type: "str"}
var CharType = VariableType{Type: "char"}
var BoolType = VariableType{Type: "bool"}
var VoidType = VariableType{Type: "void"}

type VariableType struct {
	Type        string
	SubType     *VariableType
	IsEmptyList bool
}

func (v VariableType) IsValid() bool {
	if v.Type == "List" {
		return v.SubType.IsValid()
	}

	return (v.Type == "void" ||
		v.Type == "char" ||
		v.Type == "int" ||
		v.Type == "str" ||
		v.Type == "bool")
}

func (v VariableType) ToString() string {
	if v.Type == "List" {
		return fmt.Sprintf("List<%s>", v.SubType.ToString())
	}

	return v.Type
}

func (v VariableType) Compare(t VariableType) bool {
	if v.Type == "List" && t.Type == "List" {
		if v.IsEmptyList || t.IsEmptyList {
			return true
		}

		return v.SubType.Compare(*t.SubType)
	}

	return v.Type == t.Type
}
