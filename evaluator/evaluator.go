package evaluator

import (
	"monkey2/ast"
	"monkey2/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// statements
	case *ast.Programm:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
		// expression
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}
	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}
	return result
}