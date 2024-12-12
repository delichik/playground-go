package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "D:\\Workspace\\go-test2\\getter_setter\\testdata\\model.go", nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}
	v := &TypeSpecVisitor{
		models: map[string]*model{},
	}
	ast.Walk(v, f)
	v.Gen()
}
