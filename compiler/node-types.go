package compiler

var IntType = VariableType{Type: "int"}
var StrType = VariableType{Type: "str"}
var CharType = VariableType{Type: "char"}
var BoolType = VariableType{Type: "bool"}
var VoidType = VariableType{Type: "void"}

type VariableType struct {
	Type string
}

func (v VariableType) IsValid() bool {
	return (v.Type == "void" ||
		v.Type == "char" ||
		v.Type == "int" ||
		v.Type == "str" ||
		v.Type == "bool")
}

func (v VariableType) ToString() string {
	return v.Type
}
