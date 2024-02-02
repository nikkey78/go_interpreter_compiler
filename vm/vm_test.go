package vm

import (
	"fmt"
	"monkey2/ast"
	"monkey2/compiler"
	"monkey2/lexer"
	"monkey2/object"
	"monkey2/parser"
	"testing"
)

type vmTestCase struct {
	input    string
	expected any
}

func runVmTests(t *testing.T, tests []vmTestCase) {
	t.Helper()

	for _, tt := range tests {
		program := parse(tt.input)

		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			t.Fatalf("compiler error: %s", err)
		}

		vm := New(comp.Bytecode())
		err = vm.Run()
		if err != nil {
			t.Fatalf("ev error: %s", err)
		}

		// stackElem := vm.StackTop()
		stackElem := vm.LastPoppedStackElem()

		// t.Logf("%q", comp.Bytecode())
		// t.Logf("%q", stackElem)

		testExpectedObject(t, tt.expected, stackElem)
	}
}

func parse(input string) *ast.Programm {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}

func testExpectedObject(t *testing.T, expected any, actual object.Object) {
	t.Helper()

	switch expected := expected.(type) {
	case int:
		err := testIntegerObject(int64(expected), actual)
		if err != nil {
			t.Errorf("testIntegerObject failed: %s", err)
		}
	}
}

func testIntegerObject(expected int64, actual object.Object) error {
	result, ok := actual.(*object.Integer)
	if !ok {
		return fmt.Errorf("object is not Integer. got=%T (%+v)", actual, actual)
	}
	if result.Value != expected {
		return fmt.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
	}
	return nil
}

func TestIntegerArithmetic(t *testing.T) {
	tests := []vmTestCase{
		{"1", 1},
		{"2", 2},
		{"1 + 5", 6}, // FIXME  number 5 on top of stack
	}
	runVmTests(t, tests)
}
