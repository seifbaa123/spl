package node

var Void = VariableType{Type: "void"}
var Char = VariableType{Type: "char"}
var Int = VariableType{Type: "int"}
var Bool = VariableType{Type: "bool"}

type VariableType struct {
	Type string
}
