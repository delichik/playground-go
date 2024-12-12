package visitor

import (
	"go/ast"
	"strings"
)

type RelatedModel struct {
	ast.StructType
	Targets []string
}

type RelatedVisitor struct {
	packageName string
	models      map[string]*RelatedModel
}

func NewRelatedVisitor() *RelatedVisitor {
	return &RelatedVisitor{
		models: make(map[string]*RelatedModel),
	}
}

func (v *RelatedVisitor) Clear() {
	clear(v.models)
}

func (v *RelatedVisitor) Models() map[string]*RelatedModel {
	return v.models
}

func (v *RelatedVisitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}

	switch n := node.(type) {
	case *ast.Package:
		v.packageName = n.Name
	case *ast.GenDecl:
		if n.Doc == nil {
			return v
		}
		var targets []string
		for _, line := range n.Doc.List {
			if strings.HasPrefix(strings.TrimSpace(line.Text[2:]), "rel:") {
				targets = append(targets, strings.TrimSpace(line.Text[2:])[4:])
			}
		}
		if len(targets) == 0 {
			return v
		}

		for _, spec := range n.Specs {
			switch child := spec.(type) {
			case *ast.TypeSpec:
				t, ok := child.Type.(*ast.StructType)
				if !ok {
					return nil
				}

				v.models[child.Name.Name] = &RelatedModel{
					StructType: *t,
					Targets:    targets,
				}
				return nil
			}
		}
	}

	return v
}

func (v *RelatedVisitor) PackageName() string {
	return v.packageName
}
