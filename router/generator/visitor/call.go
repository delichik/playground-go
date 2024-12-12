package visitor

import (
	"go/ast"
)

type CallVisitor struct {
	relations map[string]map[string]struct{}
}

func NewCallVisitor() *CallVisitor {
	return &CallVisitor{
		relations: make(map[string]map[string]struct{}),
	}
}

func (v *CallVisitor) Relations() map[string]map[string]struct{} {
	return v.relations
}

func (v *CallVisitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}

	switch n := node.(type) {
	case *ast.CallExpr:
		if n.Fun == nil {
			return nil
		}

		ile, ok := n.Fun.(*ast.IndexListExpr)
		if !ok {
			return v
		}
		idn := ile.X.(*ast.Ident)
		if idn.Name != "CopyContent" {
			return v
		}
		f, fm := getName(ile.Indices[0])
		t, tm := getName(ile.Indices[1])
		if f == "" || t == "" || fm || !tm {
			return v
		}
		r, ok := v.relations[f]
		if !ok {
			r = make(map[string]struct{})
			v.relations[f] = r
		}
		r[t] = struct{}{}
	}
	return v
}

func getName(n ast.Expr) (name string, model bool) {
	if n == nil {
		return
	}
	switch n := n.(type) {
	case *ast.Ident:
		return n.Name, false
	case *ast.SelectorExpr:
		idn := n.X.(*ast.Ident)
		if idn.Name != "model" {
			return
		}
		return n.Sel.Name, true
	}
	return
}
