package visitor

import (
	"go/ast"
	"strings"
)

type ModelVisitor struct {
	packageName string
	models      map[string]*ast.StructType
	prefix      string
	suffix      string
}

func NewModelVisitor() *ModelVisitor {
	return &ModelVisitor{
		models: make(map[string]*ast.StructType),
	}
}

func (v *ModelVisitor) WithPrefix(prefix string) {
	v.prefix = prefix
}

func (v *ModelVisitor) WithSuffix(suffix string) {
	v.suffix = suffix
}

func (v *ModelVisitor) Models() map[string]*ast.StructType {
	return v.models
}

func (v *ModelVisitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}

	switch n := node.(type) {
	case *ast.Package:
		v.packageName = n.Name
	case *ast.GenDecl:
		for _, spec := range n.Specs {
			switch child := spec.(type) {
			case *ast.TypeSpec:
				t, ok := child.Type.(*ast.StructType)
				if !ok {
					return nil
				}
				if !strings.HasPrefix(child.Name.Name, v.prefix) {
					continue
				}

				if !strings.HasSuffix(child.Name.Name, v.suffix) {
					continue
				}

				v.models[child.Name.Name] = t
				return nil
			}
		}
	}
	return v
}

func (v *ModelVisitor) PackageName() string {
	return v.packageName
}
