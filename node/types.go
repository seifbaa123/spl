package node

var Void = VariableType{Type: "void"}
var Int = VariableType{Type: "int"}

type VariableType struct {
	Type string
}
