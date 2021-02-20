package evaluator

import (
	"kdnakt/writing-an-interpreter-in-go/ast"
	"kdnakt/writing-an-interpreter-in-go/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}

