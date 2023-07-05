package node

var Void = VariableType{Type: "void"}
var Char = VariableType{Type: "char"}
var Int = VariableType{Type: "int"}

type VariableType struct {
	Type string
}
