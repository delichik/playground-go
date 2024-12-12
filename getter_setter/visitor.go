package main

import (
	"fmt"
	"go/ast"
	"strings"
)

type modelField struct {
	name     string
	typeName string
	r        bool
	w        bool
}

type model struct {
	name        string
	validFields []modelField
	methods     []string
}

func (m *model) gen() {
	structName := m.name
	funcs := []string{}
	for _, field := range m.validFields {
		fieldName := field.name
		fieldFUName := strings.ToUpper(fieldName[0:1]) + fieldName[1:]
		fieldTypeName := field.typeName
		if field.r {
			funcs = append(funcs, "Get"+fieldFUName+"() "+fieldTypeName)
			fmt.Println("func (m *" + structName + ") Get" + fieldFUName + "() " + fieldTypeName + " {")
			fmt.Println("\t return m." + fieldName + "")
			fmt.Println("}")
			fmt.Println()
		}
		if field.w {
			funcs = append(funcs, "Set"+fieldFUName+"(v "+fieldTypeName+")")
			fmt.Println("func (m *" + structName + ") Set" + fieldFUName + "(v " + fieldTypeName + ") {")
			fmt.Println("\tm." + fieldName + " = v")
			fmt.Println("}")
			fmt.Println()
		}
	}
	if len(funcs) > 0 || len(m.methods) > 0 {
		fmt.Println("func (m *" + structName + ") __" + structName + "Iface () {}")
		fmt.Println()

		fmt.Println("type " + structName + "Iface interface {")
		for _, f := range m.methods {
			fmt.Println("\t" + f)
		}
		if len(m.methods) > 0 {
			fmt.Println()
		}
		for _, f := range funcs {
			fmt.Println("\t" + f)
		}
		if len(funcs) > 0 {
			fmt.Println()
		}
		fmt.Println("\t__" + structName + "Iface()")
		fmt.Println("}")
		fmt.Println()
	}
}

type TypeSpecVisitor struct {
	packageName string
	models      map[string]*model
}

func (v *TypeSpecVisitor) Gen() {
	fmt.Println("package " + v.packageName)
	fmt.Println()
	for _, m := range v.models {
		m.gen()
	}
}

func (v *TypeSpecVisitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.File:
		v.packageName = n.Name.Name
	case *ast.FuncDecl:
		if n.Recv == nil || len(n.Recv.List) == 0 {
			break
		}
		recv := n.Recv.List[0]

		structName := strings.TrimPrefix(getTypeName(recv.Type), "*")
		m, ok := v.models[structName]
		if !ok {
			m = &model{
				name: structName,
			}
			v.models[structName] = m
		}

		f := n.Name.Name + "("
		if n.Type.Params != nil {
			for _, param := range n.Type.Params.List {
				for _, name := range param.Names {
					f += name.Name
					f += ","
				}
				f = f[:len(f)-1]
				f += " "
				f += getTypeName(param.Type)
			}
		}
		f += ")"

		if n.Type.Results != nil {
			if len(n.Type.Results.List) > 1 {
				f += "("
			}
			for _, param := range n.Type.Results.List {
				for _, name := range param.Names {
					f += name.Name
					f += ","
				}
				f = f[:len(f)-1]
				f += " "
				f += getTypeName(param.Type)
			}
			if len(n.Type.Results.List) > 1 {
				f += ")"
			}
		}
		m.methods = append(m.methods, f)
	case *ast.TypeSpec:
		nt, ok := n.Type.(*ast.StructType)
		if !ok {
			break
		}
		structName := n.Name.Name
		m, ok := v.models[structName]
		if !ok {
			m = &model{
				name: structName,
			}
			v.models[structName] = m
		}
		for _, param := range nt.Fields.List {
			if len(param.Names) == 0 {
				continue
			}
			f := modelField{
				name:     param.Names[0].Name,
				typeName: getTypeName(param.Type),
				r:        false,
				w:        false,
			}
			f.r, f.w = getMode(param.Tag)
			if f.r || f.w {
				m.validFields = append(m.validFields, f)
			}
		}
	}
	return v
}

func getMode(tag *ast.BasicLit) (read, write bool) {
	if tag == nil {
		return
	}
	s := tag.Value
	s = strings.ToLower(s[1 : len(s)-1])
	parts := strings.Split(s, ",")
	for _, part := range parts {
		kv := strings.Split(part, ":")
		if kv[0] != "gs" {
			continue
		}
		read = strings.Contains(kv[1], "r")
		write = strings.Contains(kv[1], "w")
	}
	return
}

func getTypeName(n ast.Expr) string {
	if n == nil {
		return ""
	}
	switch n := n.(type) {
	case *ast.Ident:
		return n.Name
	case *ast.SelectorExpr:
		idn := n.X.(*ast.Ident)
		return idn.Name + "." + n.Sel.Name
	case *ast.ArrayType:
		return "[]" + getTypeName(n.Elt)
	case *ast.StarExpr:
		return "*" + getTypeName(n.X)
	case *ast.MapType:
		return "map[" + getTypeName(n.Key) + "]" + getTypeName(n.Value)
	case *ast.StructType:
		if n.Fields != nil && len(n.Fields.List) > 0 {
			panic("only struct{} in fields of struct has been supported")
		}
		return "struct{}"
	case *ast.ChanType:
		return "chan " + getTypeName(n.Value)
	}
	return ""
}
