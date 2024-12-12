package main

import (
	"go/ast"
	"strings"
)

type TypeConfig struct {
	Wired     bool
	TableName string
}

type Model struct {
	PackageName string
	StructName  string
	Config      *TypeConfig
}

type TypeSpecVisitor struct {
	packageName    string
	lastTypeConfig *TypeConfig
	models         []Model
}

func (v *TypeSpecVisitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.File:
		v.packageName = n.Name.Name
	case *ast.Comment:
		text := strings.ToLower(strings.TrimSpace(n.Text))
		if strings.HasPrefix(text, "@dym/") {
			kv := strings.Split(text[5:], ":")
			switch kv[0] {
			case "wired":
				v.lastTypeConfig.Wired = true
			case "table":
				v.lastTypeConfig.TableName = kv[1]
			}
		}
	case *ast.TypeSpec:
		if v.lastTypeConfig.Wired {
			_, ok := n.Type.(*ast.StructType)
			if ok {
				if v.lastTypeConfig.TableName == "" {
					v.lastTypeConfig.TableName = n.Name.Name
				}
				v.models = append(v.models, Model{
					PackageName: v.packageName,
					StructName:  n.Name.Name,
					Config:      v.lastTypeConfig,
				})
			}
		}
		v.lastTypeConfig = &TypeConfig{}
	}
	return v
}
